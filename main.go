package main

import (
	"agenda/db"
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

	db.LogAllDates()

	exports.App.Static("/", "./public")

	exports.App.Get("/api", handlers.GetAllHandler)

	exports.App.Get("/api/:date", handlers.GetDateHandler)

	exports.App.Get("/api/year/:year", handlers.GetYearHandler)

	exports.App.Get("/api/:year/:month", handlers.GetYearMonthHandler)

	exports.App.Post("/api/new", handlers.PostNewDateHandler)

	exports.App.Patch("/api/update", handlers.PatchDateHandler)

	exports.App.Delete("/api/delete", handlers.DeleteDateHandler)

	exports.App.Get("*", func(c *fiber.Ctx) error {
		return c.Redirect("/")
	})

	fmt.Printf("⚡ | WebServer listening on [http://localhost%s]!\n", PORT)
	log.Fatal(exports.App.Listen(PORT))
}

/////////////////////////////////////////////////////////////////////

/*exports.App.Post("/post", func(c *fiber.Ctx) error {
	payload := struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return err
	}
	fmt.Printf(payload.Name + ": " + payload.Email + "\n")
	return c.JSON(payload)
})*/

///////////////////////////////////////////////////////////////////////
