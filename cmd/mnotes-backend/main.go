package main

import (
	"fmt"
	"os"

	handler "github.com/MichaelGit202/mnotes-backend/internal"
	"github.com/gofiber/fiber/v2"
)

func main() {
	cwd, _ := os.Getwd()
	fmt.Println("Working directory:", cwd)

	app := fiber.New()

	// Serve static files
	app.Static("/", "./web/client/mnotes/build", fiber.Static{
		Browse:   false,
		Index:    "",
		Compress: true,
	})

	// API routes
	api := app.Group("/api")
	handler.RegisterAPIRoutes(api)

	// SPA fallback
	app.Get("*", func(c *fiber.Ctx) error {
		return c.SendFile("./web/client/mnotes/build/index.html")
	})

	app.Listen(":3000")
}
