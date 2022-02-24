package main

import (
	"fmt"
	"sync"
	"time"
)

//defining variable for all to use

const conferenceTickets = 50

var conferenceName = "Go conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

//map of userdata containing different data types
type UserData struct {
	firstname  string
	lastname   string
	email      string
	userticket uint
}

var wg = sync.WaitGroup{}

func main() {

	//infinite loop is a loop whose condition is always true so we can write it as for true

	greetUsers()

	//calling userinput function with its parameters
	firstname, lastname, email, tickets := getUserInput()

	//calling of validate function
	isValidName, isValidEmail, isValidTicket := validateUserInput(firstname, lastname, email, tickets)

	if isValidName && isValidEmail && isValidTicket {

		bookTicket(tickets, firstname, lastname, email)

		//we are adding one go function in add function
		wg.Add(1)
		//go keyword helps to make the project more concurrent
		go sendTicket(tickets, firstname, lastname, email)

		//function for getting first name of the user
		//bookings = append(bookings, firstname+" "+lastname)
		firstNames := getFirstNames()
		fmt.Printf("The first name of our user in our system is:%v\n", firstNames)

		noRemainingTicket := remainingTickets == 0

		if noRemainingTicket {
			println("Sorry no more ticket is available for our conference")
			//break
		}

	} else {

		if !isValidName {
			fmt.Println("Your first and last name is too short")
		}

		if !isValidEmail {
			fmt.Println("Your email address dont have @ sign")
		}
		if !isValidTicket {
			fmt.Println("your no of ticket booking is not correct")
		}

	}
	//waits for previous function to be completed
	wg.Wait()

}

//passing conference name as string
//group of code that belongs together
func greetUsers() {
	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

//fuc for printing first names
//get is used to return some data to main function
func getFirstNames() []string {
	firstNames := []string{}

	//this loop ends when iterated over all elements in bookings list
	//underscore means that there is variable there but we are not using it
	for _, booking := range bookings {

		firstNames = append(firstNames, booking.firstname)
	}

	return firstNames

}

func getUserInput() (string, string, string, uint) {
	var firstname string
	var lastname string
	var email string
	var tickets uint

	//this function will greet the user

	fmt.Println("Enter your name:")
	fmt.Scan(&firstname)

	fmt.Println("Enter your lastname:")
	fmt.Scan(&lastname)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("Enter your purchased tickets")
	fmt.Scan(&tickets)

	return firstname, lastname, email, tickets

}

func bookTicket(tickets uint, firstname string, lastname string, email string) {
	remainingTickets = remainingTickets - tickets

	//used to access the struct
	var userData = UserData{

		firstname:  firstname,
		lastname:   lastname,
		email:      email,
		userticket: tickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for purchasing %v tickets it will be sent to your %v email\n", firstname, lastname, tickets, email)
	fmt.Printf("%v tickets are avaialable for %v conference\n", remainingTickets, conferenceName)

}

//simply using to send the email
func sendTicket(tickets uint, firstname string, lastname string, email string) {
	var ticket = fmt.Sprintf("%v tickets for %v %v", tickets, firstname, lastname)

	//something happens in 10 seconds or stops execution for 10 seconds
	time.Sleep(10 * time.Second)
	fmt.Println(".........................")
	fmt.Printf("Sending Ticket:\n %v \n to email address %v\n ", ticket, email)
	fmt.Println(".........................")

	//i am done executing so main function dont have to wait for me
	//decreases the add counter
	wg.Done()
}
