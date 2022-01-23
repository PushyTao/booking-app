package main

import (
	"booking-app/helper"
	"fmt"
	"strings"
)

var (
	bookings    []string
	firstName   string
	lastName    string
	userTickets uint
	email       string
)
var conferenceName string = "Go Conference"
var remainingTickets uint = 50

const conferenceTickets int = 50 // const can't change the value of it

func main() {
	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	for remainingTickets > 0 && len(bookings) < 50 {
		getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidUserInput(firstName, lastName, email, userTickets, remainingTickets)
		// !代表取反操作
		if isValidName && isValidEmail && isValidTicketNumber {
			bookTickets()
			// print first names func
			fmt.Printf("The firstnames of our bookings: %v\n", printFirstNames(bookings))
			var noTicketsRemaining = remainingTickets == 0
			if noTicketsRemaining {
				// end the program
				fmt.Println("Our conference is booked out. Come back next year")
				break
			}
		} else {
			incorrectInput(isValidName, isValidEmail, isValidTicketNumber)
		}
	}
}

func greetUsers(confName string, confTick int, remainTick uint) {
	fmt.Printf("Welcome to %v booking application\n", confName)
	fmt.Printf("We hava total of %v tickets and %v are still available\n", confTick, remainTick)
	fmt.Println("Get your tickets here to attend")
}

func printFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func getUserInput() {
	// ask user's name
	fmt.Println("Please enter your first name: ")
	fmt.Scan(&firstName) // 加上取地址符，否则不会输入
	fmt.Println("Please enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Please enter your email: ")
	fmt.Scan(&email)
	fmt.Println("Please enter number of tickets: ")
	fmt.Scan(&userTickets)
}

func incorrectInput(isValidName bool, isValidEmail bool, isValidTicketNumber bool) {
	if !isValidName {
		fmt.Println("the firstname of lastname is too short")
	}
	if !isValidEmail {
		fmt.Println("the email address you entered is not contain @ sign")
	}
	if !isValidTicketNumber {
		fmt.Println("number of tickets you entered is not valid")
	}
}

func bookTickets() {
	remainingTickets = remainingTickets - userTickets
	// bookings[0] = firstName + " " + lastName
	// append添加一个元素到切片的末尾，并返回一个添加后的切片，显然可以直接赋给bookings
	bookings = append(bookings, firstName+" "+lastName)
	fmt.Printf("User (%v %v) booked %v tickets.\n", firstName, lastName, userTickets)
	fmt.Printf("You will receive an confirmation email at %v\n", email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
