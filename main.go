package main

import (
	"fmt"
	"sync"
	"time"
)

const ticketsAvailable int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isVaildTicketNumber := ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isVaildTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)

			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)

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
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v of which %v are remianing\n", ticketsAvailable, remainingTickets)
	fmt.Println("Click here to book the tickets")
}

func getFirstNames() []string {
	fNames := []string{}
	for _, booking := range bookings {
		fNames = append(fNames, booking.firstName)
	}
	return fNames
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

	//creating a map
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of Bookings is %v\n", bookings)
	fmt.Printf("Thank you %v %v, for buying %v tickets.\nYou will receive confirmation email on %v.\n",
		firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var tickets = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("###################")
	fmt.Printf("Sending ticket\n %v to email address %v\n", tickets, email)
	fmt.Println("###################")
	wg.Done()
}
