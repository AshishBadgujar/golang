package main

import (
	"fmt"
	"sync"
	"time"
)

const conferenceTickets = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName   string
	lastName    string
	userTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUser()

	firstName, lastName, userTickets := getUserInput()

	isValidTicket := userTickets > 0 && userTickets <= remainingTickets

	if isValidTicket {
		bookTickets(userTickets, firstName, lastName)

		wg.Add(1)       //no of new threads adding
		go sendTicket() // go key word
		firstNames := getFnames()

		fmt.Printf("Our bookings %v \n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("Our conference is booked.")
		}
	} else {
		if !isValidTicket {
			fmt.Println("Please enter valide tickets")
		}
		fmt.Println("Your input is invalid")
	}
	wg.Wait() //wait for all threads to complete
}
func sendTicket() {
	//this function will run in saperate thread
	time.Sleep(10 * time.Second)
	fmt.Println("#################")
	fmt.Println("Sending ticket to your email...")
	fmt.Println("#################")

	wg.Done() //removes thread
}
