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
	tagRepo  *models.TagRepository
}

func NewBookService(userRepo *models.UserRepository, bookRepo *models.BookRepository, tagRepo *models.TagRepository) *BookService {
	return &BookService{
		userRepo: userRepo,
		bookRepo: bookRepo,
		tagRepo:  tagRepo,
	}
}

func (s *BookService) CreateBook(title, description string, authors []uint, tags []string) (*models.Book, error) {
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

	var Tags []*models.Tag
	for _, t := range tags {
		tag, err := s.tagRepo.GetByName(t)
		if err != nil {
			return nil, err
		}
		Tags = append(Tags, tag)
	}

	book := &models.Book{
		Title:       title,
		Description: description,
		Authors:     Authors,
		Tags:        Tags,
	}

	err = s.bookRepo.Create(book)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (s *BookService) GetRecentBooks(limit int) ([]*models.Book, error) {
	if limit < 1 {
		limit = 10
	}
	return s.bookRepo.GetRecent(limit)
}
