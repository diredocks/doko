package main

import (
	"log"

	"doko/database"
	_ "doko/docs"
	"doko/middleware"
	"doko/router"
	"doko/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

// @title Doko API
// @version 0.1
// @description Doko backend API documentation.
// @BasePath /
func main() {
	app := fiber.New(fiber.Config{
		StructValidator: &utils.StructValidator{Validator: validator.New()},
		ErrorHandler:    middleware.ErrorHandler,
	})

	app.Use(middleware.Recover())
	// app.Use(cors.New())

	database.ConnectDB()

	router.SetupRoutes(app)
	log.Fatal(app.Listen(":4096"))
}
