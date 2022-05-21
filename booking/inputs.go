package main

import (
	"fmt"
)

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


func getTicketsCount() uint {
	var ticketsCount uint;
	fmt.Println("How many tickets you want? ")
	fmt.Scanln(&ticketsCount)

	return ticketsCount
}
