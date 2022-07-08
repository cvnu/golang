package main

import "fmt"

func main() {
	var conferenceName string = "Go conference"
	const conferenceTickets int = 29
	var remainingTickets int = 5

	fmt.Println("welcome: ", conferenceName, "booking service")
	fmt.Println("Get your tickets here to attend")
	fmt.Println("Remaining tickets: ", remainingTickets, "after this we have nothing more im sorry")
	fmt.Println("how many tickets where there ? ", conferenceTickets, "is the amount that we had")

	var firstName string
	var lastName string
	var email string
	var userTickets string
	//ask user for their name
	fmt.Println("enter here your first name")
	fmt.Scan(&firstName)

	fmt.Println("enter here your last name")
	fmt.Scan(&lastName)

	fmt.Println("enter here your email")
	fmt.Scan(&email)

	fmt.Println("enter here your wanted amount of tickets here: ")
	fmt.Scan(&userTickets)

	fmt.Println("thank you for", firstName, lastName, "for booking", userTickets, "tickets. u will recieve a confirmation email at", email)

}
