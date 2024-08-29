package helper

import "booking-app/models"

func GetFirstName(bookings []models.UserData) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var first = booking.FirstName
		firstNames = append(firstNames, first)
	}

	return firstNames
}
