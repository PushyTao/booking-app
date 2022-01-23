package helper

import (
	"fmt"
	"strings"
)

func ValidUserInput(first string, last string, email string, userTick uint, remaTick uint) (bool, bool, bool) {
	isValidName := len(first) >= 2 && len(last) >= 2
	isValidEmail := strings.Contains(email, "@") // 字符串是否含有'@'
	isValidTicketNumber := userTick > 0 && userTick <= remaTick
	return isValidName, isValidEmail, isValidTicketNumber
}
func IncorrectInput(isValidName bool, isValidEmail bool, isValidTicketNumber bool) {
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
func GreetUsers(confName string, confTick int, remainTick uint) {
	fmt.Printf("Welcome to %v booking application\n", confName)
	fmt.Printf("We hava total of %v tickets and %v are still available\n", confTick, remainTick)
	fmt.Println("Get your tickets here to attend")
}
func PrintFirstNames(bookings []map[string]string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		// var names = strings.Fields(booking)
		// case 1
		// var firname = booking["firstName"]
		// firstNames = append(firstNames, firname)
		// case 2
		firstNames = append(firstNames, booking["firstName"])
	}
	return firstNames
}
