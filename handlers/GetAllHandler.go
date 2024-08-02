package handlers

import (
	"agenda/db"

	"github.com/gofiber/fiber/v2"
)

func GetAllHandler(c *fiber.Ctx) error {
	dates, err := db.GetAllDates()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(dates)
}
