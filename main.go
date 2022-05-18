package main

import (
	"fmt"
)

func main() {
	// Say hello and introduction
	const conferenceName string = "Go Conference"
	var availableTickets uint = 1400
	var remainingTickets uint = availableTickets

	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v ticket and %v are still available\n", availableTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	// Array of clients 
	clients := []string{}

	for (remainingTickets > 0) {
		// Ask users for ther informations
		var (
			username string
			email string
			age uint
			ticketsCount uint
		)

		fmt.Println("Enter your full name: ")
		fmt.Scanln(&username)
		fmt.Println("Enter your email: ")
		fmt.Scanln(&email)
		fmt.Println("How old are you? ")
		fmt.Scanln(&age)
		
		if (age >= 18) {
			// Book
			fmt.Println("How many tickets you want? ")
			fmt.Scanln(&ticketsCount)
			
			// Inscription message 
			if (ticketsCount > 4 ) {
				fmt.Println("NOTE! >> You can only book 4 tickets as maximum")
			} else {
				// Calculate the remaining tickets
				remainingTickets = remainingTickets - ticketsCount;
				// Successfully booking tickets message
				fmt.Printf("Hello %v. You have booked %v tickets, Thank you\n", username, ticketsCount)
				fmt.Printf("Remaining tickets:  %v.\n", remainingTickets)
				// Add client to the database
				clients = append(clients, username)
				// Print out all the clients
				for _, client := range clients {
					fmt.Println(client)
				}
			}
			fmt.Printf("Clients: %v\n", clients)
		} else {
			fmt.Println("Sorry, you'r still young for booking")
		}
	}
}