package db

import (
	"agenda/exports"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func GetYear(year string) ([]exports.DateType, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"date": bson.M{"$regex": fmt.Sprintf("^%s", year)}}
	cur, err := exports.GetDatesColl().Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var dates []exports.DateType
	for cur.Next(ctx) {
		var date exports.DateType
		err := cur.Decode(&date)
		if err != nil {
			return nil, err
		}
		dates = append(dates, date)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return dates, nil
}
