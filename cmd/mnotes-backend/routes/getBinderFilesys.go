package routes

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

func GetBinderFolders(c *fiber.Ctx) error {
	binderName := c.Params("binderName")
	dirPath := filepath.Join("binders", binderName)

	files, err := os.ReadDir(dirPath)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Binder not found",
		})
	}

	folderNames := []string{}
	for _, file := range files {
		if file.IsDir() {
			folderNames = append(folderNames, file.Name())
		}
	}

	return c.JSON(folderNames)
}

func GetPage(c *fiber.Ctx) error {
	binderName := c.Params("binderName")
	folderName := c.Params("folderName")
	pageName := c.Params("page")

	filePath := filepath.Join("binders", binderName, folderName, pageName+".json")

	data, err := os.ReadFile(filePath)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Page not found",
		})
	}

	// Optionally validate JSON
	var parsed interface{}
	if err := json.Unmarshal(data, &parsed); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Invalid JSON in page",
		})
	}

	return c.JSON(parsed)
}

func GetBinderThumbnails(c *fiber.Ctx) error {
	filePath := filepath.Join("binders", "index.json")

	data, err := os.ReadFile(filePath)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "index.json not found",
		})
	}

	// Optionally validate it's valid JSON
	var parsed interface{}
	if err := json.Unmarshal(data, &parsed); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "invalid JSON format in index.json",
		})
	}

	return c.JSON(parsed)
}
