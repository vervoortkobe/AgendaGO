package handlers

import (
	"agenda/db"
	"agenda/exports"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func PostNewDateHandler(c *fiber.Ctx) error {
	var date exports.DateType
	if err := c.BodyParser(&date); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	exists, err := db.CheckDateExists(date.Date)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if exists {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "Date already exists"})
	}

	if err := db.InsertDate(date); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	fmt.Print(date)
	return c.Status(fiber.StatusCreated).JSON(date)
}
