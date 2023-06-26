package main

import (
	"fmt"
	"strings"
)

const ticketsAvailable int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = []string{}

func main() {

	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isVaildTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isVaildTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)

			fNames := getFirstNames()
			fmt.Printf("These are first Names of bookings: %v\n", fNames)

			ticketsSoldout := remainingTickets <= 0

			if ticketsSoldout {
				fmt.Println("Sorry! We are booked out. Try again in the next year")
				break
			}

		} else {
			if !isVaildTicketNumber {
				fmt.Printf("Sorry! Only %v tickets are left\n", remainingTickets)
			}
			if !isValidEmail {
				fmt.Println("Email is Invalid as it does not contain @")
			}
			if !isValidName {
				fmt.Println("Your First Name or Last Name is too short")
			}
			continue
		}
	}

}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v of which %v are remianing\n", ticketsAvailable, remainingTickets)
	fmt.Println("Click here to book the tickets")
}

func getFirstNames() []string {
	fNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		var firstName = names[0]
		fNames = append(fNames, firstName)
	}
	return fNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isVaildTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isVaildTicketNumber
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter First Name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter Last Name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter email: ")
	fmt.Scan(&email)

	fmt.Println("Enter tickets required: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)
	fmt.Printf("Thank you %v %v, for buying %v tickets.\nYou will receive confirmation email on %v.\n",
		firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}
