// Package handlers
package handlers

import (
	"strconv"

	"doko/services"

	"github.com/gofiber/fiber/v3"
)

type CreateBookRequest struct {
	Title       string   `json:"title" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Authors     []uint   `json:"authors" validate:"required"`
	Tags        []string `json:"tags"`
}

type CreateBookResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type RecentBookRequest struct {
	Limit int `uri:"limit"`
}

type RecentBookResponse struct {
	ID          string       `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Authors     []AuthorInfo `json:"authors"`
	Tags        []string     `json:"tags"`
}

type AuthorInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Error on create book request",
			"data":    nil,
		})
	}

	book, err := bh.bookService.CreateBook(in.Title, in.Description, in.Authors, in.Tags)
	if err != nil {
		if err == services.ErrBookExisted {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"status":  "error",
				"message": "Book already existed",
				"data":    nil,
			})
		}
		if err == services.ErrMissingAuthors {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":  "error",
				"message": "Missing authors",
				"data":    nil,
			})
		}
		if err == services.ErrAuthorNotFound {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":  "error",
				"message": "Author not found",
				"data":    nil,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Error on creating book",
			"data":    nil,
		})
	}

	newBook := CreateBookResponse{
		ID:          strconv.FormatUint(uint64(book.ID), 10),
		Title:       book.Title,
		Description: book.Description,
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Success create book",
		"data":    newBook,
	})
}

func (bh *BookHandler) GetRecentBooks(c fiber.Ctx) error {
	var in RecentBookRequest
	if err := c.Bind().URI(&in); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Param limit should be an integer >= 1",
			"data":    nil,
		})
	}

	books, err := bh.bookService.GetRecentBooks(in.Limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Error on fetching recent books",
			"data":    nil,
		})
	}

	var out []RecentBookResponse
	for _, b := range books {
		authors := []AuthorInfo{}
		for _, a := range b.Authors {
			authors = append(authors, AuthorInfo{
				ID:   strconv.FormatUint(uint64(a.ID), 10),
				Name: a.Username,
			})
		}
		tags := []string{}
		for _, t := range b.Tags {
			tags = append(tags, t.Name)
		}
		out = append(out, RecentBookResponse{
			ID:          strconv.FormatUint(uint64(b.ID), 10),
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
