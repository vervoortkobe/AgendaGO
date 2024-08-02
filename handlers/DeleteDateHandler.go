package handlers

import (
	"agenda/db"

	"github.com/gofiber/fiber/v2"
)

func DeleteDateHandler(c *fiber.Ctx) error {
	dateParam := c.Params("date")
	exists, err := db.CheckDateExists(dateParam)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if exists {
		result, err := db.DeleteDate(dateParam)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		if result {
			return c.SendStatus(fiber.StatusNoContent)
		} else {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "Failed to delete the date"})
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Date doesn't exist"})
}
