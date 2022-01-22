package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50 // const can't change the value of it
	var remainingTickets uint = 50
	// %T可以用来输出类型
	// fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We hava total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		//array
		// var bookings [50]string //"tao", "tao2", "tao3"
		// bookings[0] = "tao"     // the index is 0->n-1
		// bookings[1] = "tao2"    // one by one
		// slice
		var bookings []string
		var firstName string
		var lastName string
		var userTickets uint
		var email string
		// ask user's name
		fmt.Println("Please enter your first name: ")
		fmt.Scan(&firstName) // 加上取地址符，否则不会输入
		fmt.Println("Please enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Println("Please enter your email: ")
		fmt.Scan(&email)
		fmt.Println("Please enter number of tickets: ")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets
		// bookings[0] = firstName + " " + lastName
		// append添加一个元素到切片的末尾，并返回一个添加后的切片，显然可以直接赋给bookings
		bookings = append(bookings, firstName+" "+lastName)
		/*
			fmt.Printf("the whole slice : %v\n", bookings)
			fmt.Printf("The first value is %v\n", bookings[0])
			fmt.Printf("slice type is %T\n", bookings)
			fmt.Printf("slice length is :%v\n", len(bookings))
		*/
		fmt.Printf("User %v %v booked %v tickets.\n", firstName, lastName, userTickets)
		fmt.Printf("You will receive an confirmation email at %v\n", email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("The firstnames of our bookings: %v\n", firstNames)

	}
}
