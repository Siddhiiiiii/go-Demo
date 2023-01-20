package main

import (
	"booing-app/helper"
	"fmt"
)

var confName = "Go Conference"

const conftickets = 50

var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName   string
	lastName    string
	email       string
	userTickets uint
}

func main() {

	greetusers()

	for remainingTickets != 0 {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)

			firstNames := getfirstnames()

			fmt.Printf("These are the list of %v first Names\n", firstNames)

			var res bool = remainingTickets == 0

			if res {
				fmt.Printf("Tickets sold out!!\n")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Name is wrong")
			}
			if !isValidEmail {
				fmt.Println("Email is wrong")
			}
			if !isValidTicketNumber {
				fmt.Println("Tickets are wrong")
			}
		}
	}
}

func greetusers() {
	fmt.Println("Welcome to", confName, "booking application")
	fmt.Println("We have total of", conftickets, "tickets and", remainingTickets, "are still left!")
	fmt.Println("Get your tickets from here to attend")
}

func getfirstnames() []string {
	firstNames := []string{} //Slices

	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:   firstName,
		lastName:    lastName,
		email:       email,
		userTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings %v\n", bookings)

	fmt.Printf("Thank You for %v %v for booking the tickets . You will get the congirmation mail at %v\n", firstName, lastName, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, confName)

}
