package db

import (
	"agenda/exports"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateDate(date exports.DateType) error {

	var existingDate exports.DateType
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
	if date.HourlyData.Hour != existingDate.HourlyData.Hour {
		updateFields = append(updateFields, bson.E{Key: "hourlyData.hour", Value: date.HourlyData.Hour})
	}
	if date.HourlyData.Data != existingDate.HourlyData.Data {
		updateFields = append(updateFields, bson.E{Key: "hourlyData.data", Value: date.HourlyData.Data})
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
