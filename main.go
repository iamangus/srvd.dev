package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	// Initialize template engine
	engine := html.New("./views", ".html")

	// Create Fiber app with HTML template engine
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Serve static files
	app.Static("/", "./public")

	// Define routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "SRVD.DEV - Your Development Partner",
		})
	})

	// Start server
	log.Fatal(app.Listen(":3000"))
}
