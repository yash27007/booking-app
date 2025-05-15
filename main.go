package main

import (
	"booking-app/helper"
	"fmt"
)

const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings []UserData

type UserData struct {
	firstName  string
	lastName   string
	email      string
	numTickets uint
}

func main() {
	helper.GreetUsers(conferenceName, conferenceTickets, remainingTickets)

	for remainingTickets > 0 {
		// Get user input
		var firstName, lastName, email string
		var userTickets uint

		fmt.Print("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Print("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Print("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Print("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		// Validate input
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(firstName, lastName, email, userTickets)

			firstNames := getFirstNames()
			fmt.Printf("Current bookings: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("All tickets are sold out! Come back next year.")
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name is too short.")
			}
			if !isValidEmail {
				fmt.Println("Email address is invalid.")
			}
			if !isValidTicketNumber {
				fmt.Printf("We only have %v tickets remaining. Cannot book %v tickets.\n", remainingTickets, userTickets)
			}
		}
	}
}

func bookTicket(firstName, lastName, email string, userTickets uint) {
	remainingTickets -= userTickets

	userData := UserData{
		firstName:  firstName,
		lastName:   lastName,
		email:      email,
		numTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("Thank you %v %v for booking %v tickets. A confirmation email has been sent to %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v.\n\n", remainingTickets, conferenceName)
}

func getFirstNames() []string {
	var firstNames []string
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}
