package handlers

import (
	"agenda/db"
	"agenda/exports"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func PostNewDateHandler(c *fiber.Ctx) error {
	var date exports.Appointment
	if err := c.BodyParser(&date); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	// Generate a new UUID for the appointment
	date.Id = uuid.New().String()

	// Parse the incoming date string
	incomingDateStr := date.Date
	fmt.Printf("Received date: %s\n", incomingDateStr)
	parsedDate, err := time.Parse("01/02/2006", incomingDateStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid date format"})
	}

	// Convert to YYYY-MM-DD format
	formattedDate := parsedDate.Format("2006-01-02")
	fmt.Printf("Formatted date: %s\n", formattedDate)

	// Replace the date in the `date` object with the formatted date
	date.Date = formattedDate

	// Check if the date already exists in the database
	exists, err := db.CheckAppointmentExists(date.Id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if exists {
		fmt.Printf("Date %s already exists\n", date.Date)
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "Date already exists"})
	}

	// Insert the date into the database
	if err := db.InsertDate(date); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	fmt.Printf("Date inserted successfully: %v\n", date)
	return c.Status(fiber.StatusCreated).JSON(date)
}
