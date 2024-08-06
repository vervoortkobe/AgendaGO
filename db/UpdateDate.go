package db

import (
	"agenda/exports"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateDate(date exports.Appointment) error {

	var existingDate exports.Appointment
	err := exports.GetDatesColl().FindOne(context.TODO(), bson.D{{Key: "date", Value: date.Date}}).Decode(&existingDate)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("no date found: %s", date.Date)
		}
		return err
	}

	updateFields := bson.D{}
	if date.Date != existingDate.Date {
		updateFields = append(updateFields, bson.E{Key: "date", Value: date.Date})
	}
	if date.Hour != existingDate.Hour {
		updateFields = append(updateFields, bson.E{Key: "hour", Value: date.Hour})
	}
	if date.Desc != existingDate.Desc {
		updateFields = append(updateFields, bson.E{Key: "desc", Value: date.Desc})
	}

	if len(updateFields) == 0 {
		return nil
	}

	result, err := exports.GetDatesColl().UpdateOne(context.TODO(), bson.D{{Key: "date", Value: date.Date}}, bson.D{{Key: "$set", Value: updateFields}})
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("no date found: %s", date.Date)
	}

	return nil
}
