package main

import (
	"booking-app/validation"
	"fmt"
)

// Create a new structure
type ClientData struct {
	username string
	email string
	age uint
	ticketsCount uint
}

// Create an Empty List of clients
var clients = make([]ClientData, 0)


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

				// create a map for clients
				var clientData = ClientData {
					username: username,
					email: email,
					age: age,
					ticketsCount: ticketsCount,
				}
			
				clients = append(clients, clientData)

				// for _, client := range clients {
				// 	fmt.Println(client.username)
				// }

				// Append client to the console
				appendClients(clients)
			}
		} else {
			fmt.Println("Sorry, you'r not allowed for booking")
		}
	}
}

func appendClients(clients []ClientData) {
	fmt.Printf("Clients: %v\n", clients)
}