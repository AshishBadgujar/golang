package main

import (
	"fmt"
)

func getUserInput() (string, string, uint) {
	var firstName string
	var lastName string
	var userTickets uint

	fmt.Print("Enter First Name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter last Name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter Tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, userTickets
}
func greetUser() {
	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets NOW!")
}

func getFnames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}
func bookTickets(userTickets uint, firstName string, lastName string) {
	remainingTickets = remainingTickets - userTickets
	var userData = UserData{
		firstName:   firstName,
		lastName:    lastName,
		userTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("Thank you %v booking %v tickets\n", firstName, userTickets)
	fmt.Printf("%v tickets remaining\n", remainingTickets)
}
