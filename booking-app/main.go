package main

import "fmt"

func main(){
const conferenceTickets = 50
var remainingTickets = 50 
var conferenceName = "GO Con"
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

fmt.Printf("User %v booked %v tickets for the show.\n" ,userName ,userTickets)
}
