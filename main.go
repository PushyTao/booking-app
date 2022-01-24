package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

var (
	firstName   string
	lastName    string
	userTickets uint
	email       string
)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

// var bookings = make([]map[string]string, 0) // 初始长度 == 0
var bookings = make([]UserData, 0) // 初始长度 == 0
var conferenceName string = "Go Conference"
var remainingTickets uint = 50

const conferenceTickets int = 50 // const can't change the value of it

var waitgrp = sync.WaitGroup{}

func main() {
	helper.GreetUsers(conferenceName, conferenceTickets, remainingTickets)

	// for remainingTickets > 0 && len(bookings) < 50 {
	getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := helper.ValidUserInput(firstName, lastName, email, userTickets, remainingTickets)
	// !代表取反操作
	if isValidName && isValidEmail && isValidTicketNumber {
		bookTickets()
		//go关键字，开启一个新的线程
		waitgrp.Add(1)
		go sendTickets(userTickets, firstName, lastName, email)
		// print first names func
		fmt.Printf("The firstnames of our bookings: %v\n", PrintFirstNames(bookings))
		var noTicketsRemaining = remainingTickets == 0
		if noTicketsRemaining {
			// end the program
			fmt.Println("Our conference is booked out. Come back next year")
			// break
		}
	} else {
		helper.IncorrectInput(isValidName, isValidEmail, isValidTicketNumber)
	}
	// }
	waitgrp.Wait()
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

func bookTickets() {
	remainingTickets = remainingTickets - userTickets
	// create a map for user
	/*
		var userData = make(map[string]string)
		userData["firstName"] = firstName
		userData["lastName"] = lastName
		userData["email"] = email
		userData["number"] = strconv.FormatUint(uint64(userTickets), 10)
	*/
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	// bookings[0] = firstName + " " + lastName
	// append添加一个元素到切片的末尾，并返回一个添加后的切片，显然可以直接赋给bookings
	bookings = append(bookings, userData)
	fmt.Printf("User (%v %v) booked %v tickets.\n", firstName, lastName, userTickets)
	fmt.Printf("You will receive an confirmation email at %v\n", email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
	fmt.Printf("The list of all booking: %v\n", bookings)
}
func PrintFirstNames(bookings []UserData) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		// var names = strings.Fields(booking)
		// case 1
		// var firname = booking["firstName"]
		// firstNames = append(firstNames, firname)
		// case 2
		// firstNames = append(firstNames, booking["firstName"])
		// case 3
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func sendTickets(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(5 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for (%v %v)", userTickets, firstName, lastName)
	fmt.Println("$$$$$$$$$$$$$$$$$$$$$$$")
	fmt.Printf("Sending tickets:\n %v \n to email address %v\n", ticket, email)
	fmt.Println("$$$$$$$$$$$$$$$$$$$$$$$")
	waitgrp.Done()
}
