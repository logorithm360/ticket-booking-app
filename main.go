package main

import "fmt"

func main() {
	// ticket booking app in Golang

	var conference = "Conference"
	const tickets int = 50
	var availableTickets = 50
	fmt.Println("Welcome to cromeon", conference, "ticket booking ")
	fmt.Println("We have total of", tickets, "and tickets remained are", availableTickets)
	fmt.Println("Get your tickets here to attend")

}
