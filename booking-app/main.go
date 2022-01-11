package main

import "fmt"

func main() {
	var remainingTickets = 50 // or var remainingTickets int = 50
	// This explicit mention is more appropriate to use for special use cases like here it should be var remainingTickets uint = 50
	// The "uint" is unsigned integer and would only take 0 and +ve values, as this variable is being used for remainingTickets and the value for it cannot be negative.
	var conferenceName = "GO Con" // or conferenceName := "Go Con" , In this case Golang automatically understands and it is called syntacticSugar in programming.
	//Note this cannot be done in the below case of constants. and also if the syntacticSugar i.e. := is used the explicit mention of "int" or "uint" cannot be used
	const conferenceTickets = 50
	fmt.Println("-------Checking Variable Types-------")
	fmt.Printf("Type for conferenceTickers is %T, for remainingTickets is %T, for conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	fmt.Println("--------------")
	fmt.Println("Using Println --> Welcome to", conferenceName, "booking application")
	fmt.Printf("Using Printf --> Welcome to %v booking application\n", conferenceName)
	fmt.Println("Using Println --> We have total of", conferenceTickets, "tickets and", remainingTickets, "are still avaiable")
	fmt.Printf("Using Printf --> We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend now!!")

	var userName string
	var userTickets int

	userName = "Tom"
	userTickets = 2

	fmt.Printf("User %v booked %v tickets for the show.\n", userName, userTickets)
}
