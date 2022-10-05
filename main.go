package main

import "fmt"

func main() {
	var confername = "Rayyan Darbar!!"
	const capacity = 50
	var available = 50
	var booked = 0
	fmt.Printf("Welcome to %v, Conference \n", confername)
	fmt.Printf("Conference Capacity %v, avaialbe %v booked %v : \n", capacity, available, booked)

	var username string
	var tickets int
	var emailId string

	fmt.Println("Please enter name :")
	fmt.Scan(&username)
	fmt.Println("Please enter emaild :")
	fmt.Scan(&emailId)
	fmt.Println("How many tickets:")
	fmt.Scan(&tickets)

	fmt.Printf("Hello %v, Booking has confirmed!!-%v", username, tickets)

}
