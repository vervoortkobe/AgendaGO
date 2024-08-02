package db

import (
	"agenda/exports"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CheckDateExists(date string) (bool, error) {
	filter := bson.M{"date": date}
	var result exports.DateType
	err := exports.GetDatesColl().FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
