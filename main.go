package main

import (
	"fmt"
	"regexp"
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

		if (isValidClient(username,email, age)) {
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

func sayHello(subject string, availableTickets uint, remainingTickets uint) {
	fmt.Printf("Welcome to our %v booking application 👋\n", subject)
	fmt.Printf("We have total of %v ticket and %v are still available\n", availableTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend 🎟️")
}

func getClientInfo() (string, string, uint) {
	var (
		username string
		email string
		age uint
	)
	// Get client main information
	fmt.Println("Enter your full name: ")
	fmt.Scanln(&username)
	fmt.Println("Enter your email: ")
	fmt.Scanln(&email)
	fmt.Println("How old are you? ")
	fmt.Scanln(&age)

	return username, email, age
}

func isValidClient(username string, email string, age uint) bool {
	if (len(username) >= 3 && isEmailValid(email) && age >= 18) {
		return true
	}
	return false
}

func isEmailValid(email string) bool {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
    return emailRegex.MatchString(email)
}


func getTicketsCount() uint {
	var ticketsCount uint;
	fmt.Println("How many tickets you want? ")
	fmt.Scanln(&ticketsCount)

	return ticketsCount
}

func calculateRemainingTickets(ticketsCount uint, remainingTickets uint) uint {
	return remainingTickets - ticketsCount
}

func congrateClient(username string, ticketsCount uint) {
	fmt.Printf("Hello %v. You have booked %v tickets, Thank you\n", username, ticketsCount)
}

func appendClients(clients []string, email string) {
	clients = append(clients, email)
	fmt.Printf("Clients: %v\n", clients)
}