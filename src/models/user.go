package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string     `gorm:"uniqueIndex;not null" json:"username"`
	Email     string     `gorm:"uniqueIndex;not null" json:"email"`
	Password  string     `gorm:"not null" json:"-"`
	Nickname  string     `json:"nickname"`
	LastLogin *time.Time `json:"last_login"`
	Books     []*Book    `gorm:"many2many:book_authors;" json:"books"`
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(email, username, passwordHash string) (*User, error) {
	user := &User{
		Email:    email,
		Username: username,
		Password: passwordHash,
	}

	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) GetByEmail(email string) (*User, error) {
	var user User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetByID(id uint) (*User, error) {
	var user User
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetByIDs(ids []uint) ([]*User, error) {
	users := []*User{}
	if err := r.db.Where("id IN ?", ids).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) GetByUsername(username string) (*User, error) {
	var user User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Delete(id uint) error {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var user User
	if err := tx.First(&user, "id = ?", id).Error; err != nil {
		tx.Rollback()
		return err
	}

	var books []*Book
	if err := tx.Model(&user).Association("Books").Find(&books); err != nil {
		tx.Rollback()
		return err
	}

	for _, book := range books {
		var authors []*User
		if err := tx.Model(book).Association("Authors").Find(&authors); err != nil {
			tx.Rollback()
			return err
		}
		if len(authors) <= 1 {
			if err := tx.Delete(book).Error; err != nil {
				tx.Rollback()
				return err
			}
		} else {
			if err := tx.Model(book).Association("Authors").Delete(&user); err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	if err := tx.Delete(&user).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *UserRepository) Update(id uint, updateUser User) (*User, error) {
	var user User
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	user.Email = updateUser.Email
	user.Username = updateUser.Username
	user.Password = updateUser.Password
	user.Nickname = updateUser.Nickname

	if err := r.db.Save(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
