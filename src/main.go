package main

import (
	"log"

	"doko/database"
	"doko/router"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()
	// app.Use(cors.New())

	database.ConnectDB()

	router.SetupRoutes(app)
	log.Fatal(app.Listen(":4096"))
}
