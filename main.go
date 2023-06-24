package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const ticketsAvailable = 50
	remainingTickets := 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v of which %v are remianing\n", ticketsAvailable, remainingTickets)
	fmt.Println("Click here to book the tickets")

	bookings := []string{}

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets int

		fmt.Println("Enter First Name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter Last Name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter email: ")
		fmt.Scan(&email)

		fmt.Println("Enter tickets required: ")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isVaildTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		bookings = append(bookings, firstName+" "+lastName)

		if isValidName && isValidEmail && isVaildTicketNumber {
			remainingTickets = remainingTickets - userTickets
			fmt.Printf("Thank you %v %v, for buying %v tickets.\nYou will receive confirmation email on %v.\n",
				firstName, lastName, userTickets, email)

			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			fNames := []string{}

			for _, booking := range bookings {
				var names = strings.Fields(booking)
				var firstName = names[0]
				fNames = append(fNames, firstName)
			}

			ticketsSoldout := remainingTickets <= 0

			if ticketsSoldout {
				fmt.Println("Sorry! We are booked out. Try again in the next year")
				break
			}

			fmt.Printf("These are first Names of bookings: %v\n", fNames)
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
