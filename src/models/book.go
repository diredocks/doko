// Package models
package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title       string     `gorm:"not null" json:"title"`
	Description string     `gorm:"not null" json:"description"`
	Authors     []*User    `gorm:"many2many:book_authors;" json:"authors"`
	Tags        []*Tag     `gorm:"many2many:book_tags;" json:"tags"`
	Chapters    []*Chapter `gorm:"foreignKey:BookID;constraint:OnDelete:CASCADE;" json:"chapters"`
}

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (r *BookRepository) base() *gorm.DB {
	return r.db.Model(&Book{}).Preload("Authors").Preload("Tags").Preload("Chapters")
}

func (r *BookRepository) Create(book *Book) error {
	return r.db.Create(book).Error
}

func (r *BookRepository) GetAll() ([]*Book, error) {
	var books []*Book
	err := r.base().Find(&books).Error
	return books, err
}

func (r *BookRepository) GetByID(id uint) (*Book, error) {
	var book Book
	err := r.base().First(&book, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *BookRepository) GetRecent(limit int) ([]*Book, error) {
	var books []*Book
	err := r.base().Order("created_at DESC").Limit(limit).Find(&books).Error
	return books, err
}

func (r *BookRepository) Find(title string, authors []uint, tags []uint) ([]*Book, error) {
	var books []*Book

	db := r.base()

	if title != "" {
		db = db.Where("title = ?", title)
	}

	// matches any of the authors
	if len(authors) > 0 {
		sub := r.db.
			Table("book_authors").
			Select("book_id").
			Where("user_id IN ?", authors)

		db = db.Where("id IN (?)", sub)
	}

	// matches all of the tags
	if len(tags) > 0 {
		sub := r.db.
			Table("book_tags").
			Select("book_id").
			Where("tag_id IN ?", tags).
			Group("book_id").
			Having("COUNT(DISTINCT tag_id) = ?", len(tags))

		db = db.Where("books.id IN (?)", sub)
	}

	if err := db.Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

func (r *BookRepository) Delete(id uint) (int64, error) {
	book := Book{Model: gorm.Model{ID: id}}

	tx := r.db.Begin()

	if err := tx.Model(&book).Association("Authors").Clear(); err != nil {
		tx.Rollback()
		return 0, err
	}
	if err := tx.Model(&book).Association("Tags").Clear(); err != nil {
		tx.Rollback()
		return 0, err
	}

	if err := tx.Where("book_id = ?", id).Delete(&Chapter{}).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	result := tx.Delete(&book)
	if result.Error != nil {
		tx.Rollback()
		return 0, result.Error
	}

	return result.RowsAffected, tx.Commit().Error
}
