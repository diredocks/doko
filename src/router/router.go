// Package router
package router

import (
	"doko/database"
	"doko/handlers"
	"doko/models"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	userRepo := models.NewUserRepository(database.DB)
	userHandler := handlers.NewUserHandler(userRepo)
	user := api.Group("/user")
	user.Post("/register", userHandler.Register)
}
