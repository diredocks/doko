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

func (r *UserRepository) GetByUsername(username string) (*User, error) {
	var user User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Delete(id uint) error {
	var user User
	if err := r.db.Where("id = ?", id).Delete(&user).Error; err != nil {
		return err
	}
	return nil
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
