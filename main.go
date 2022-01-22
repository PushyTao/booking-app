package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50 // const can't change the value of it
	var remainingTickets = 50
	// %T可以用来输出类型
	// fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We hava total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var userName string
	var userTickets int
	// ask user's name
	userName = "Tom"
	userTickets = 2

	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)

}
