package handlers

import (
	"agenda/db"

	"github.com/gofiber/fiber/v2"
)

func GetAppointmentHandler(c *fiber.Ctx) error {
	appointmentParam := c.Params("id")
	appointment, err := db.GetAppointment(appointmentParam)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(appointment)
}
