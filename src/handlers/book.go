// Package handlers
package handlers

import (
	"strconv"

	"doko/services"

	"github.com/gofiber/fiber/v3"
)

type CreateBookRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Authors     []uint `json:"authors"`
}

type CreateBookResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
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

	book, err := bh.bookService.CreateBook(in.Title, in.Description, in.Authors)
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
