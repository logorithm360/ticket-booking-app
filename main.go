package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	// ticket booking app in Golang

	/*
		PRINT FORMATTING CONSTRUCTS
		%v	the value in a default format
		when printing structs, the plus flag (%+v) adds field names

		%#v	a Go-syntax representation of the value
			(floating-point infinities and NaNs print as ±Inf and NaN)

		%T	a Go-syntax representation of the type of the value

		%%	a literal percent sign; consumes no value

	*/

	var conference = "Conference"
	var availableTickets uint = 50
	var bookings []string   // Move outside the loop to accumulate all bookings
	var firstNames []string // Move outside the loop to accumulate all first names
	fmt.Printf("Welcome to cromeon %v ticket booking \n", conference)
	fmt.Printf("We have total of %v and tickets remained are %v \n", availableTickets, availableTickets)
	fmt.Println("Get your tickets here to attend")

	// the power of infinite loop
	for {
		// The only error here is: the loop continues to ask endless times while tickets are only 50
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter your first name: ")
		// this scan function takes the inputed value and assigns it to a variable pointed to which is firstName
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email name: ")
		fmt.Scan(&email)

		fmt.Println("How many tickets do you want?")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicket := userTickets > 0 && userTickets < availableTickets

		// managing the number of tickets a user needs buy
		if userTickets > availableTickets {
			fmt.Printf("Sorry we only have %v tickets\n", availableTickets)
			continue
		}

		if isValidName && isValidEmail && isValidTicket {

			availableTickets = availableTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			// logic for calling user's first name
			firstNames = nil // Reset the slice each time to avoid duplicates
			// for each loop in golang
			for _, booking := range bookings {
				names := strings.Fields(booking)
				name := names[0]
				firstNames = append(firstNames, name)
			}

			fmt.Printf("Thank you %v %v for booking %v tickets.\n We will confirm and inform you at %v \n\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v AVAILABLE FOR THE %v \n", availableTickets, strings.ToUpper(conference))

			// retrieving values from a slice is the same as in arrays
			// fmt.Printf("%v booked the ticket \n\n", bookings[0])
			fmt.Printf("NAMES THAT BOOKED TICKETS ARE:\n %v \n\n", strings.Join(firstNames, " \n"))

			if availableTickets == 0 {
				fmt.Printf("Thank you %v %v for booking %v tickets.\n We will confirm and inform you at %v \n\n", firstName, lastName, userTickets, email)
				time.Sleep(5 * time.Second)

				fmt.Printf("Thank you for your request.\n WE ARE OUT OF TICKETS TILL TOMORROW!\n")
				break

			}
		} else {
			if !isValidName {
				fmt.Println("name is too short, please enter a long name")
			}
			if !isValidEmail {
				fmt.Println("email doesn't contain @, please enter a valid name")

			}
			if !isValidTicket {
				fmt.Println("Please enter a valid number of tickets following the remaining tickets")

			}
		}

	}

}
