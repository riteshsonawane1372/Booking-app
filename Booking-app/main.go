// Everything in Go in a package 

package main

import	"fmt" // A basic Go package 

func main() {

	var conferenceTickets =50
	var remainingTickets =50
	var FirstName=""
	var LastName=""
	var requiredTickets uint =0

	fmt.Print("Welcome to Go Conference")
	fmt.Println("\n \t\t\t \nTicket for Go conference")
	fmt.Printf("\n Remaining Tickets: %v ",remainingTickets)
	fmt.Printf("\n Out Off : %v",conferenceTickets)
	fmt.Print("\n \n")
	fmt.Print("Enter Your First Name :")
	fmt.Scan(&FirstName)
	fmt.Print("Enter Your last Name :")
	fmt.Scan(&LastName)
	fmt.Printf("Every Good Afternoon %v %v Welcome to Go conference",FirstName,LastName)
	fmt.Printf("\n Remaining Tickets %v kindly book soon ",remainingTickets)
	fmt.Printf(" \n Enter no. Ticket :")
	fmt.Scan(&requiredTickets)
	remainingTickets=remainingTickets-requiredTickets
	fmt.Printf("Now Remaining Tickets are %v",remainingTickets)

}