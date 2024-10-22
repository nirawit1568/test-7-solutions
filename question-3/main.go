package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nirawit1568/test-7-solutions/handlers"
)

func main() {

	// Route
	app := fiber.New()
	app.Get("/beef/summary", handlers.BeefSummary)
	app.Listen(":3000")
}
