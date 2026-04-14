package main

import (
	"log"

	"doko/database"
	"doko/middleware"
	"doko/router"
	"doko/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

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
