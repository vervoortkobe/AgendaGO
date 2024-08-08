package handlers

import (
	"agenda/db"

	"github.com/gofiber/fiber/v2"
)

func GetMonthHandler(c *fiber.Ctx) error {
	yearParam := c.Params("year")
	monthParam := c.Params("month")
	yearMonth, err := db.GetMonth(yearParam, monthParam)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(yearMonth)
}
