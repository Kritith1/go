package main

import "strings"

func validateUserInput(firstname string, lastname string, email string, tickets uint) (bool, bool, bool) {
	isValidName := len(firstname) >= 2 && len(lastname) >= 2

	//validates if email contains @ or not
	isValidEmail := strings.Contains(email, "@")

	//for user ticket
	isValidTicket := tickets > 0 && tickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicket

}
