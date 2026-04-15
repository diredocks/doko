// Package router
package router

import (
	"doko/database"
	"doko/handlers"
	"doko/models"
	"doko/services"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	userRepo := models.NewUserRepository(database.DB)
	bookRepo := models.NewBookRepository(database.DB)
	tagRepo := models.NewTagRepository(database.DB)
	chapterRepo := models.NewChapterRepository(database.DB)

	bookService := services.NewBookService(userRepo, bookRepo, tagRepo, chapterRepo)

	userHandler := handlers.NewUserHandler(userRepo)
	bookHandler := handlers.NewBookHandler(bookService)
	tagHandler := handlers.NewTagHandler(tagRepo)

	user := api.Group("/user")
	user.Post("/register", userHandler.Register)
	user.Delete("/:id", userHandler.Delete)

	book := api.Group("/book")
	book.Post("/", bookHandler.CreateBook)
	book.Get("/:id", bookHandler.GetBook)
	book.Delete("/:id", bookHandler.DeleteBook)
	book.Post("/:id/chapter", bookHandler.CreateChapter)
	book.Get("/recent/:limit", bookHandler.GetRecentBooks)

	tag := api.Group("/tag")
	tag.Get("/", tagHandler.GetAllTags)
}
