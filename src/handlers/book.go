// Package handlers
package handlers

import (
	"errors"

	"doko/middleware"
	"doko/services"

	"github.com/gofiber/fiber/v3"
)

type CreateBookRequest struct {
	Title       string   `json:"title" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Authors     []uint   `json:"authors" validate:"required,min=1"`
	Tags        []string `json:"tags"`
}

type CreateBookResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type RecentBookRequest struct {
	Limit int `uri:"limit,default:10" validate:"min=1"`
}

type RecentBookResponse struct {
	ID          uint     `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Authors     []string `json:"authors"`
	Tags        []string `json:"tags"`
}

type CreateChapterRequest struct {
	BookID  uint   `json:"book_id" validate:"required"`
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
	Order   int    `json:"order" validate:"required"`
}

type CreateChapterResponse struct {
	ID      uint   `json:"id"`
	BookID  uint   `json:"book_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Order   int    `json:"order"`
}

type BookHandler struct {
	bookService *services.BookService
}

func NewBookHandler(bookService *services.BookService) *BookHandler {
	return &BookHandler{
		bookService: bookService,
	}
}

func (bh *BookHandler) CreateBook(c fiber.Ctx) error {
	var in CreateBookRequest
	if err := c.Bind().Body(&in); err != nil {
		return middleware.NewValidationError(err)
	}

	book, err := bh.bookService.CreateBook(in.Title, in.Description, in.Authors, in.Tags)
	if err != nil {
		if errors.Is(err, services.ErrBookExisted) {
			return middleware.NewAppError(fiber.StatusConflict, "Book already existed", err)
		}
		if errors.Is(err, services.ErrAuthorNotFound) {
			return middleware.NewAppError(fiber.StatusBadRequest, "Author not found", err)
		}
		return middleware.NewAppError(fiber.StatusInternalServerError, "Error on creating book", err)
	}

	newBook := CreateBookResponse{
		ID:          book.ID,
		Title:       book.Title,
		Description: book.Description,
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Success create book",
		"data":    newBook,
	})
}

func (bh *BookHandler) CreateChapter(c fiber.Ctx) error {
	var in CreateChapterRequest
	if err := c.Bind().Body(&in); err != nil {
		return middleware.NewValidationError(err)
	}

	chapter, err := bh.bookService.CreateChapter(in.BookID, in.Title, in.Content, in.Order)
	if err != nil {
		if errors.Is(err, services.ErrBookNotFound) {
			return middleware.NewAppError(fiber.StatusNotFound, "Book not found", err)
		}
		return middleware.NewAppError(fiber.StatusInternalServerError, "Error on creating chapter", err)
	}

	out := CreateChapterResponse{
		ID:      chapter.ID,
		BookID:  chapter.BookID,
		Title:   chapter.Title,
		Content: chapter.Content,
		Order:   chapter.Order,
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Success create chapter",
		"data":    out,
	})
}

func (bh *BookHandler) DeleteBook(c fiber.Ctx) error {
	var in DeleteBookRequest
	if err := c.Bind().URI(&in); err != nil {
		return middleware.NewValidationError(err)
	}

	if err := bh.bookService.DeleteBook(in.ID); err != nil {
		if errors.Is(err, services.ErrBookNotFound) {
			return middleware.NewAppError(fiber.StatusNotFound, "Book not found", err)
		}
		return middleware.NewAppError(fiber.StatusInternalServerError, "Error on deleting book", err)
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Success delete book",
	})
}

func (bh *BookHandler) GetRecentBooks(c fiber.Ctx) error {
	var in RecentBookRequest
	if err := c.Bind().URI(&in); err != nil {
		return middleware.NewValidationError(err)
	}

	books, err := bh.bookService.GetRecentBooks(in.Limit)
	if err != nil {
		return middleware.NewAppError(fiber.StatusInternalServerError, "Error on fetching recent books", err)
	}

	out := []RecentBookResponse{}
	for _, b := range books {
		authors := []string{}
		for _, a := range b.Authors {
			authors = append(authors, a.Username)
		}
		tags := []string{}
		for _, t := range b.Tags {
			tags = append(tags, t.Name)
		}
		out = append(out, RecentBookResponse{
			ID:          b.ID,
			Title:       b.Title,
			Description: b.Description,
			Authors:     authors,
			Tags:        tags,
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Success fetch recent books",
		"data":    out,
	})
}
