package main

import (
	"booking-app/helper"
	"fmt"
	"strconv"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTikcets int = 50
	var remainingTickets uint = 50

	type UserData struct {
	  firstName string 
	  lastName string
	  email string
	  numberOfTickets uint
	}

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTikcets, remainingTickets)
	fmt.Printf("Get your tickets here to attend\n")

	var bookings = make([]UserData, 0)

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

    var userData UserData		
    userData.firstName = firstName
    userData.lastName = lastName
    userData.email = email
    userData.numberOfTickets = userTickets

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

