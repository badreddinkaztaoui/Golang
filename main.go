package main

import (
	"booking-app/validation"
	"fmt"
	"strconv"
)

// Array of clients
var clients = make([]map[string]string, 0)

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
				var clientData = make(map[string]string)
				clientData["username"] = username
				clientData["email"] = email
				clientData["age"] = strconv.FormatUint(uint64(age), 10)
				clientData["ticketsCount"] = strconv.FormatUint(uint64(ticketsCount), 10)

				// Append client to the console
				clients = append(clients, clientData)
				appendClients(clients)
			}
		} else {
			fmt.Println("Sorry, you'r not allowed for booking")
		}
	}
}

func appendClients(clients []map[string]string) {
	fmt.Printf("Clients: %v\n", clients)
}