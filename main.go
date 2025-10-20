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

}
