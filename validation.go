package main

import (
	"regexp"
)

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
