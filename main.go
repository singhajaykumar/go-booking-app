package main

import "fmt"

func main() {

	fmt.Println("Welcome to Booking application of GO Conference")

	var conferenceName string = "Go Conference"
	var noTicket int

	fmt.Println(conferenceName)

	fmt.Println("Enter how many tickets you need ?")
	fmt.Scan(&noTicket)

	fmt.Printf("You booked ticket = %v", noTicket)

}
