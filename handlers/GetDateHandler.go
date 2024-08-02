package handlers

import (
	"agenda/dbactions"

	"github.com/gofiber/fiber/v2"
)

func GetDateHandler(c *fiber.Ctx) error {
	dateParam := c.Params("date")
	date, err := dbactions.GetDate(dateParam)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(date)
}
