package main

import (
	"fmt"
	"strings"
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
	fmt.Printf("Welcome to cromeon %v ticket booking \n", conference)
	fmt.Printf("We have total of %v and tickets remained are %v \n", availableTickets, availableTickets)
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	var bookings []string

	fmt.Println("Enter your first name: ")
	// this scan function takes the inputed value and assigns it to a variable pointed to which is firstName
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email name: ")
	fmt.Scan(&email)

	fmt.Println("How many tickets do you want?")
	fmt.Scan(&userTickets)

	availableTickets = availableTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets.\n We will confirm and inform you at %v \n\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v AVAILABLE FOR THE %v \n", availableTickets, strings.ToUpper(conference))

	// retrieving values from a slice is the same as in arrays
	// fmt.Printf("%v booked the ticket \n\n", bookings[0])
	fmt.Printf("NAMES THAT BOOKED TICKETS ARE:\n %v \n", strings.Join(bookings, " "))

}
