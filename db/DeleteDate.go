package db

import (
	"agenda/exports"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func DeleteDate(dateParam string) (bool, error) {
	result, err := exports.GetDatesColl().DeleteOne(context.TODO(), bson.D{{Key: "date", Value: dateParam}})

	if err != nil {
		return false, err
	}

	if result.DeletedCount == 0 {
		return false, nil
	}

	return true, nil
}
