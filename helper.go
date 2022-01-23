package main

import "strings"

func validUserInput(first string, last string, email string, userTick uint, remaTick uint) (bool, bool, bool) {
	isValidName := len(first) >= 2 && len(last) >= 2
	isValidEmail := strings.Contains(email, "@") // 字符串是否含有'@'
	isValidTicketNumber := userTick > 0 && userTick <= remaTick
	return isValidName, isValidEmail, isValidTicketNumber
}
