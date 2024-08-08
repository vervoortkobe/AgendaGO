package handlers

import (
	"agenda/db"

	"github.com/gofiber/fiber/v2"
)

func GetDateHandler(c *fiber.Ctx) error {
	dateParam := c.Params("date")
	yearMonth, err := db.GetDate(dateParam)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(yearMonth)
}
