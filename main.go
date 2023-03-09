package main

import (
	"fmt"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("conferenceTickets is %T, remainTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v Booking Application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get you tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)

	fmt.Println("Enter your Last name")
	fmt.Scan(&lastName)

	fmt.Println("Enter email address")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets you would like purchase")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you %v %v for booking tickets. You will receive an email at %v \n", firstName, lastName, email)
	fmt.Printf("%v Remaining Tickets for %v\n", remainingTickets, conferenceName)
}
