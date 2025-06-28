package main

import (
	"github.com/MichaelGit202/mnotes-backend/handler"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Static("/", "./client/mnotes/build")

	// Register API routes
	api := app.Group("/api")
	handler.RegisterAPIRoutes(api)

	app.Get("*", func(c *fiber.Ctx) error {
		return c.SendFile("./client/mnotes/build/index.html")
	})

	app.Listen(":3000")
}
