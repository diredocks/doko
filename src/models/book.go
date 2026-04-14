// Package models
package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title       string  `gorm:"not null" json:"title"`
	Description string  `gorm:"not null" json:"description"`
	Authors     []*User `gorm:"many2many:book_authors;" json:"authors"`
}

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (r *BookRepository) base() *gorm.DB {
	return r.db.Model(&Book{}).Preload("Authors")
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

func (r *BookRepository) Find(title string, authors []uint) ([]*Book, error) {
	var books []*Book

	db := r.base()

	if title != "" {
		db = db.Where("title = ?", title)
	}

	if len(authors) > 0 {
		db = db.Joins("JOIN book_authors ba ON ba.book_id = books.id")
		db = db.Where("ba.user_id IN ?", authors)
	}

	db = db.Distinct()

	if err := db.Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

func (r *BookRepository) Delete(id string) error {
	return r.db.Delete(&Book{}, "id = ?", id).Error
}
