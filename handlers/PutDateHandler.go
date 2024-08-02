package handlers

import (
	"agenda/db"
	"agenda/exports"

	"github.com/gofiber/fiber/v2"
)

func PutDateHandler(c *fiber.Ctx) error {
	var date exports.DateType
	if err := c.BodyParser(&date); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	exists, err := db.CheckDateExists(date.Date)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Date doesn't exist"})
	}

	if err := db.ReplaceDate(date); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(date)
}
