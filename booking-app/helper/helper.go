package helper

func GetFirstName(bookings []map[string]string) []string {
	firstNames := []string{}
	for _, name := range bookings {
		var first = name["firstName"]
		firstNames = append(firstNames, first)
	}

	return firstNames
}
