package dbactions

import (
	"fmt"
	"log"
)

func LogAllDates() {
	dates, err := GetAllDates()
	if err != nil {
		log.Fatalf("Error retrieving dates: %v", err)
	}

	fmt.Println("All dates in the database:")
	for _, date := range dates {
		fmt.Printf("Date: %v, HourlyData: %v\n", date.Date, date.HourlyData)
	}
}
