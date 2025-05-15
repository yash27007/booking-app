package helper

import (
	"fmt"
	"strings"
)

// GreetUsers prints a welcome message to the user
func GreetUsers(confName string, totalTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application!\n", confName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", totalTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend!")
}

// ValidateUserInput checks if user input is valid
func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}
