package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetUrlHandler(c *fiber.Ctx) error {
	year := c.Params("year")
	month := c.Params("month")
	if _, err := strconv.Atoi(year); err == nil {
		if _, err := strconv.Atoi(month); err == nil {
			return c.SendFile("./public/index.html")
		}
	}
	return c.Redirect("/")
}
