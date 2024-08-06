package db

import (
	"agenda/exports"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CheckAppointmentExists(id string) (bool, error) {
	filter := bson.M{"id": id}
	var result exports.Appointment
	err := exports.GetDatesColl().FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
