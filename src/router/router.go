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
	userHandler := handlers.NewUserHandler(userRepo)
	user := api.Group("/user")
	user.Post("/register", userHandler.Register)
	user.Delete("/:id", userHandler.Delete)

	bookRepo := models.NewBookRepository(database.DB)
	tagRepo := models.NewTagRepository(database.DB)
	bookService := services.NewBookService(userRepo, bookRepo, tagRepo)
	bookHandler := handlers.NewBookHandler(bookService)
	book := api.Group("/book")
	book.Post("/", bookHandler.CreateBook)
	book.Get("/recent", bookHandler.GetRecentBooks)
}
