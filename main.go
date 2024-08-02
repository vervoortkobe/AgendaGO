package main

import (
	"agenda/dbactions"
	"agenda/exports"
	"agenda/handlers"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")

	exports.GetMongoClient()
	fmt.Print("💽 | Connected to MongoDB!\n")

	dbactions.LogAllDates()

	exports.App.Static("/", "./public")

	exports.App.Get("/api", handlers.GetAllHandler)

	exports.App.Get("/api/:date", handlers.GetDateHandler)

	/*app.Post("/add-date", func(c *fiber.Ctx) error {
		var date Date
		if err := c.BodyParser(&date); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
		}

		exists, err := checkDateExists(collection, date.Date)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		if exists {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "Date already exists"})
		}

		if err := insertDate(collection, date); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return c.Status(fiber.StatusCreated).JSON(date)
	})*/

	/////////////////////////////////////////////////////////////////////

	exports.App.Get("/t", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	exports.App.Post("/post", func(c *fiber.Ctx) error {
		payload := struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		}{}

		if err := c.BodyParser(&payload); err != nil {
			return err
		}
		fmt.Printf(payload.Name + ": " + payload.Email + "\n")
		return c.JSON(payload)
	})

	/*exports.App.Post("/register", handlers.RegisterHandler)
	exports.App.Post("/auth", handlers.AuthHandler)

	exports.App.Post("/upload", handlers.UploadHandler)

	exports.App.Get("/:value", handlers.ImageHostBuilder)*/

	///////////////////////////////////////////////////////////////////////

	exports.App.Get("*", func(c *fiber.Ctx) error {
		return c.Redirect("/")
	})

	///////////////////////////////////////////////////////////////////////

	fmt.Printf("⚡ | WebServer listening on [http://localhost%s]!\n", PORT)
	log.Fatal(exports.App.Listen(PORT))
}
