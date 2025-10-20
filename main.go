package main

import "fmt"

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
	const tickets int = 50
	var availableTickets = 50
	fmt.Printf("Welcome to cromeon %v ticket booking \n", conference)
	fmt.Printf("We have total of %v and tickets remained are %v \n", tickets, availableTickets)
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Println("Enter your first name: ")
	// this scan function takes the inputed value and assigns it to a variable pointed to which is firstName
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email name: ")
	fmt.Scan(&email)

	fmt.Println("How many tickets do you want?")
	fmt.Scan(&userTickets)

	fmt.Printf("Thank you %v %v for booking %v tickets.\n We will confirm and inform you at %v \n\n", firstName, lastName, userTickets, email)

}
