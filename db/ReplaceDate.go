package db

import (
	"agenda/exports"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func ReplaceDate(date exports.DateType) error {
	replacement := bson.D{
		{Key: "date", Value: date.Date},
		{Key: "hour", Value: date.Hour},
		{Key: "desc", Value: date.Desc},
	}

	result, err := exports.GetDatesColl().ReplaceOne(context.TODO(), bson.D{{Key: "date", Value: date.Date}}, replacement)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("no date found: %s", date.Date)
	}

	return nil
}
