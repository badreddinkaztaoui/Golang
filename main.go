package main

import (
	"booking-app/validation"
	"fmt"
)

func main() {
	// Say hello and introduction
	const conferenceName string = "Go Conference"
	var availableTickets uint = 1400
	var remainingTickets uint = availableTickets

	sayHello(conferenceName, availableTickets, remainingTickets)

	// Array of clients 
	clients := []string{}

	for (remainingTickets > 0) {
		// Get the client informations
		username, email, age := getClientInfo()

		if (validation.IsValidClient(username,email, age)) {
			// Ask client for how mus tickets he want to book (maximum: 4)
			ticketsCount := getTicketsCount()
			
			if (ticketsCount > 4 ) {
				fmt.Println("NOTE! >> You can only book 4 tickets as maximum")
			} else {
				remainingTickets = calculateRemainingTickets(ticketsCount, remainingTickets)
				congrateClient(username, ticketsCount)
				appendClients(clients, email)
			}
		} else {
			fmt.Println("Sorry, you'r not allowed for booking")
		}
	}
}

func appendClients(clients []string, email string) {
	clients = append(clients, email)
	fmt.Printf("Clients: %v\n", clients)
}