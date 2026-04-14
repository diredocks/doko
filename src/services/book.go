// Package services
package services

import (
	"errors"

	"doko/models"

	"gorm.io/gorm"
)

var (
	ErrBookExisted    = errors.New("book already existed")
	ErrMissingAuthors = errors.New("missing authors")
	ErrAuthorNotFound = errors.New("author not found")
)

type BookService struct {
	userRepo *models.UserRepository
	bookRepo *models.BookRepository
}

func NewBookService(userRepo *models.UserRepository, bookRepo *models.BookRepository) *BookService {
	return &BookService{
		userRepo: userRepo,
		bookRepo: bookRepo,
	}
}

func (s *BookService) CreateBook(title, description string, authors []uint) (*models.Book, error) {
	books, err := s.bookRepo.Find(title, authors)
	if err != nil {
		return nil, err
	}
	if len(books) > 0 {
		return nil, ErrBookExisted
	}

	if len(authors) < 1 {
		return nil, ErrMissingAuthors
	}
	var Authors []*models.User
	for _, a := range authors {
		author, err := s.userRepo.GetByID(a)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, ErrAuthorNotFound
			}
			return nil, err
		}
		Authors = append(Authors, author)
	}

	book := &models.Book{
		Title:       title,
		Description: description,
		Authors:     Authors,
	}

	err = s.bookRepo.Create(book)
	if err != nil {
		return nil, err
	}
	return book, nil
}
