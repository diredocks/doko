// Package services
package services

import (
	"errors"
	"strings"

	"doko/models"

	"gorm.io/gorm"
)

var (
	ErrBookExisted     = errors.New("book already existed")
	ErrAuthorNotFound  = errors.New("author not found")
	ErrBookNotFound    = errors.New("book not found")
	ErrChapterNotFound = errors.New("chapter not found")
)

type BookService struct {
	userRepo    *models.UserRepository
	bookRepo    *models.BookRepository
	tagRepo     *models.TagRepository
	chapterRepo *models.ChapterRepository
}

func NewBookService(userRepo *models.UserRepository, bookRepo *models.BookRepository, tagRepo *models.TagRepository, chapterRepo *models.ChapterRepository) *BookService {
	return &BookService{
		userRepo:    userRepo,
		bookRepo:    bookRepo,
		tagRepo:     tagRepo,
		chapterRepo: chapterRepo,
	}
}

func (s *BookService) CreateBook(title, description string, authors []uint, tags []string) (*models.Book, error) {
	Authors, err := s.userRepo.GetByIDs(authors)
	if err != nil {
		return nil, err
	}
	if len(Authors) < 1 {
		return nil, ErrAuthorNotFound
	}

	books, err := s.bookRepo.Find(title, authors, []uint{})
	if err != nil {
		return nil, err
	}
	if len(books) > 0 {
		return nil, ErrBookExisted
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
	return s.bookRepo.GetRecent(limit)
}

func (s *BookService) GetBook(id uint) (*models.Book, error) {
	book, err := s.bookRepo.GetByID(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrBookNotFound
	}
	return book, err
}

func (s *BookService) DeleteBook(id uint) error {
	n, err := s.bookRepo.Delete(id)
	if err != nil {
		return err
	}
	if n == 0 {
		return ErrBookNotFound
	}
	return nil
}

func (s *BookService) GetChapter(id uint) (*models.Chapter, error) {
	chapter, err := s.chapterRepo.GetByID(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrChapterNotFound
	}
	return chapter, err
}

func (s *BookService) CreateChapter(bookID uint, title, content string, order int) (*models.Chapter, error) {
	chapter := &models.Chapter{
		BookID:  bookID,
		Title:   title,
		Content: content,
		Order:   order,
	}

	if err := s.chapterRepo.Create(chapter); err != nil {
		if strings.Contains(err.Error(), "FOREIGN KEY") {
			return nil, ErrBookNotFound
		}
		return nil, err
	}
	return chapter, nil
}
