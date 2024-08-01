package dbactions

import (
	"agenda/exports"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func GetAllDates() ([]exports.DateType, error) {
	cursor, err := exports.GetDatesColl().Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var dates []exports.DateType
	for cursor.Next(context.Background()) {
		var date exports.DateType
		if err := cursor.Decode(&date); err != nil {
			return nil, err
		}
		dates = append(dates, date)
	}
	return dates, cursor.Err()
}
