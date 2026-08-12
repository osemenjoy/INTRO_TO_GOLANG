package main

import (
	"fmt"
	"sync"
	"time"
)

var conferenceName = "Go Conference"
const conferenceTickets = 50
var remainingTickets uint = 50
var bookings = make([]userData, 0)

type userData struct {
	firstName string
	lastName string
	email string
	userTickets uint
}

var wg = sync.WaitGroup{} // wg stands for waitgroup

func main() {


	greetUser()


	firstName, lastName, email, userTickets := getuserInput()
	isValidName, isValidEmail, isUserTicketsValid := ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isUserTicketsValid {

		// book tickets
		bookTickets(firstName, lastName, userTickets ,email)

		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		// print first name function
		firstNames := getFirstNames()

		fmt.Printf("These are all the first names of the users in the application: %v\n", firstNames)
	}	else {
		if !isValidName {
			fmt.Println("First name or last name entered must be greater than 2")
		}
		if !isValidEmail {
			fmt.Println("The email entered is not a valid email, no @ sign")
		}
		if !isUserTicketsValid {
			fmt.Println("The number of ticket entered is invalid")
		}
	}
	wg.Wait()
}



func greetUser() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("There is a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}


func getuserInput() (string, string, string, uint) {
		var firstName string
		var lastName string
		var email string
		var userTickets uint


		fmt.Print("Enter your first name: ")
		fmt.Scan(&firstName)
		fmt.Print("Enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Print("Enter your email: ")
		fmt.Scan(&email)	
		fmt.Print("Enter the number of tickets to be booked: ")
		fmt.Scan(&userTickets)

		return firstName, lastName, email, userTickets
}

func bookTickets(firstName string, lastName string, userTickets uint, email string) {
		remainingTickets -= userTickets

		// create a map of the user data
		var userData = userData {
			firstName: firstName,
			lastName: lastName,
			email: email,
			userTickets: userTickets,
		}		

		// saving user details in slice
		bookings = append(bookings, userData)
		fmt.Printf("List of bookings\n %v\n",bookings)

		fmt.Printf("Thank you %v %v, you have successfully booked %v tickets. The tickets would be delivered to %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("The remaining ticket is %v\n", remainingTickets)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)

	fmt.Println("#####################")
	fmt.Printf("Sending ticket:\n %v to email address %v\n", ticket, email)
	fmt.Println("#####################")

	wg.Done()
}
