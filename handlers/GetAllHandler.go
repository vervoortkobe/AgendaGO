package handlers

import (
	"agenda/dbactions"

	"github.com/gofiber/fiber/v2"
)

func GetAllHandler(c *fiber.Ctx) error {
	dates, err := dbactions.GetAllDates()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(dates)
}
