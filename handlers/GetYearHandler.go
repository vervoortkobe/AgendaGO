package handlers

import (
	"agenda/db"

	"github.com/gofiber/fiber/v2"
)

func GetYearHandler(c *fiber.Ctx) error {
	yearParam := c.Params("year")
	year, err := db.GetYear(yearParam)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(year)
}
