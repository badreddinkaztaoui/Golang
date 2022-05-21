package main

import (
	"fmt"
)

func sayHello(subject string, availableTickets uint, remainingTickets uint) {
	fmt.Printf("Welcome to our %v booking application 👋\n", subject)
	fmt.Printf("We have total of %v ticket and %v are still available\n", availableTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend 🎟️")
}

func congrateClient(username string, ticketsCount uint) {
	fmt.Printf("Hello %v. You have booked %v tickets, Thank you\n", username, ticketsCount)
}