package routes

import (
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

func GetBinderPages(c *fiber.Ctx) error {
	binderName := c.Params("binderName")
	dirPath := filepath.Join("binders", binderName)

	files, err := os.ReadDir(dirPath)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Binder not found",
		})
	}

	pageNames := []string{}
	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".json" {
			pageNames = append(pageNames, file.Name())
		}
	}

	return c.JSON(pageNames)
}
