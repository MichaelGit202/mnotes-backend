package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Load port from env, default to 3000
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Initialize Fiber app with basic middleware
	app := fiber.New()
	app.Use(logger.New())
	app.Use(compress.New())

	// Serve static files from the React build folder
	app.Static("/", "./web/client/mnotes/build")

	// Register API routes (assumes you've defined this function)

	// Fallback: serve index.html for SPA routes (React deep links)
	app.Use(func(c *fiber.Ctx) error {
		path := c.Path()

		// Prevent fallback for known static asset prefixes
		if c.Method() == fiber.MethodGet &&
			!strings.HasPrefix(path, "/api") &&
			!strings.HasPrefix(path, "/static") &&
			!strings.HasPrefix(path, "/favicon") &&
			!strings.HasPrefix(path, "/manifest") &&
			!strings.HasPrefix(path, "/logo") {
			return c.SendFile("./web/client/mnotes/build/index.html")
		}

		// Let static files (like .js/.css) be served normally
		return c.Next()
	})

	// Start the server
	fmt.Println("updated http://localhost:" + port)
	if err := app.Listen(":" + port); err != nil {
		panic(err)
	}
}
