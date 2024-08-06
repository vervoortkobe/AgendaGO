package db

import (
	"agenda/exports"
	"context"
)

func InsertDate(date exports.Appointment) error {
	_, err := exports.GetDatesColl().InsertOne(context.Background(), date)
	return err
}
