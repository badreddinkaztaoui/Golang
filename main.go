package main

import (
	"booking-app/validation"
	"fmt"
)

// Array of clients
var clients []string

func main() {
	// Say hello and introduction
	const conferenceName string = "Go Conference"
	var availableTickets uint = 1400
	var remainingTickets uint = availableTickets

	sayHello(conferenceName, availableTickets, remainingTickets)


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
				clients = append(clients, email)
				appendClients(clients)
			}
		} else {
			fmt.Println("Sorry, you'r not allowed for booking")
		}
	}
}

func appendClients(clients []string) {
	fmt.Printf("Clients: %v\n", clients)
}