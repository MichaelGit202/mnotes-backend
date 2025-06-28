// internal/handler/api.go
package handler

import "github.com/gofiber/fiber/v2"

func RegisterAPIRoutes(app fiber.Router) {
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "pong",
		})
	})
}
