package db

import (
	"agenda/exports"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAppointment(appointmentParam string) (exports.Appointment, error) {
	var appointment exports.Appointment
	err := exports.GetDatesColl().FindOne(context.TODO(), bson.D{{Key: "id", Value: appointmentParam}}).Decode(&appointment)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return appointment, nil
		}
		return appointment, err
	}
	return appointment, nil
}
