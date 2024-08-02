package exports

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var App *fiber.App = fiber.New()

type DateType struct {
	ID         primitive.ObjectID `bson:"_id" json:"_id"`
	Date       string             `json:"date"`
	HourlyData HourlyData         `json:"hourlyData"`
}

type HourlyData struct {
	Hour int    `json:"hour"`
	Data string `json:"data"`
}

////////////////////////////////////////////////////////////

type Image struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Data      string `json:"data"`
	Timestamp int    `json:"timestamp"`
}

var EmptyImage Image = Image{}

type UserCreds struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
