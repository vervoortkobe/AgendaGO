package dbactions

import (
	"agenda/exports"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetDate(dateParam string) (exports.DateType, error) {
	var date exports.DateType
	err := exports.GetDatesColl().FindOne(context.TODO(), bson.D{{Key: "date", Value: dateParam}}).Decode(&date)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return date, nil
		}
		return date, err
	}
	return date, nil
}
