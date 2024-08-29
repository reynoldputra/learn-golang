package main

import (
	"booking-app/helper"
  "booking-app/models"
	"fmt"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTikcets int = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTikcets, remainingTickets)
	fmt.Printf("Get your tickets here to attend\n")

	bookings := make([]models.UserData, 0)

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Printf("Enter your first name: ")
		fmt.Scan(&firstName)
		fmt.Printf("Enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Printf("Enter your email: ")
		fmt.Scan(&email)
		fmt.Printf("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		var userData models.UserData
		userData.FirstName = firstName
		userData.LastName = lastName
		userData.Email = email
		userData.NumberOfTickets = userTickets

		if userTickets > remainingTickets {
			fmt.Printf("We only have %v tickets\n", remainingTickets)
			continue
		}
		bookings = append(bookings, userData)
		remainingTickets = remainingTickets - userTickets

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("Remaining tickets %v \n", remainingTickets)

		firstNames := helper.GetFirstName(bookings)

		fmt.Println(firstNames)
		if remainingTickets == 0 {
			fmt.Println("Ticket sold out")
			break
		}
	}
}
