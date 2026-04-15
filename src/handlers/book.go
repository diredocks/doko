// Package handlers
package handlers

import (
	"errors"

	"doko/middleware"
	"doko/models"
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

type DeleteBookRequest struct {
	ID uint `uri:"id" validate:"required"`
}

type RecentBookRequest struct {
	Limit int `uri:"limit,default:10" validate:"min=1"`
}

type GetBookRequest struct {
	ID uint `uri:"id" validate:"required"`
}

type ChapterSummary struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Order int    `json:"order"`
}

type BookResponse struct {
	ID          uint             `json:"id"`
	Title       string           `json:"title"`
	Description string           `json:"description"`
	Authors     []string         `json:"authors"`
	Tags        []string         `json:"tags"`
	Chapters    []ChapterSummary `json:"chapters"`
}

type CreateChapterRequestURI struct {
	BookID uint `uri:"id" validate:"required"`
}

type CreateChapterRequestBody struct {
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
	var uri CreateChapterRequestURI
	var body CreateChapterRequestBody
	if err := c.Bind().Body(&body); err != nil {
		return middleware.NewValidationError(err)
	}
	if err := c.Bind().URI(&uri); err != nil {
		return middleware.NewValidationError(err)
	}

	chapter, err := bh.bookService.CreateChapter(uri.BookID, body.Title, body.Content, body.Order)
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

func toBookResponse(book *models.Book) *BookResponse {
	authors := make([]string, 0, len(book.Authors))
	for _, a := range book.Authors {
		authors = append(authors, a.Username)
	}
	tags := make([]string, 0, len(book.Tags))
	for _, t := range book.Tags {
		tags = append(tags, t.Name)
	}
	chapters := make([]ChapterSummary, 0, len(book.Chapters))
	for _, ch := range book.Chapters {
		chapters = append(chapters, ChapterSummary{
			ID:    ch.ID,
			Title: ch.Title,
			Order: ch.Order,
		})
	}
	return &BookResponse{
		ID:          book.ID,
		Title:       book.Title,
		Description: book.Description,
		Authors:     authors,
		Tags:        tags,
		Chapters:    chapters,
	}
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

	out := []*BookResponse{}
	for _, b := range books {
		out = append(out, toBookResponse(b))
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Success fetch recent books",
		"data":    out,
	})
}

func (bh *BookHandler) GetBook(c fiber.Ctx) error {
	var in GetBookRequest
	if err := c.Bind().URI(&in); err != nil {
		return middleware.NewValidationError(err)
	}

	book, err := bh.bookService.GetBook(in.ID)
	if err != nil {
		if errors.Is(err, services.ErrBookNotFound) {
			return middleware.NewAppError(fiber.StatusNotFound, "Book not found", err)
		}
		return middleware.NewAppError(fiber.StatusInternalServerError, "Error on fetching book", err)
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Success fetch book",
		"data":    toBookResponse(book),
	})
}
