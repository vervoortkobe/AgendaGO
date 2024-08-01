package dbactions

import (
	"agenda/exports"
	"context"
)

func InsertDate(date exports.DateType) error {
	_, err := exports.GetDatesColl().InsertOne(context.Background(), date)
	return err
}
