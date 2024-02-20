package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func isUppercase(r rune) bool {
	return unicode.IsUpper(r)
}

func isLowercase(r rune) bool {
	return unicode.IsLower(r)
}

func isDigit(r rune) bool {
	return unicode.IsDigit(r)
}

func isSpecialChar(r rune) bool {
	return !unicode.IsLetter(r) && !unicode.IsDigit(r)
}

func isPasswordComplex(password string) bool {
	// Check minimum length
	if len(password) < 8 {
		return false
	}

	// Check for uppercase, lowercase, digit, and special character
	var (
		hasUppercase   = false
		hasLowercase   = false
		hasDigit       = false
		hasSpecialChar = false
	)

	for _, char := range password {
		hasUppercase = hasUppercase || isUppercase(char)
		hasLowercase = hasLowercase || isLowercase(char)
		hasDigit = hasDigit || isDigit(char)
		hasSpecialChar = hasSpecialChar || isSpecialChar(char)
	}

	return hasUppercase && hasLowercase && hasDigit && hasSpecialChar
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	if isPasswordComplex(password) {
		fmt.Println("Password is complex and meets the criteria.")
	} else {
		fmt.Println("Password does not meet the complexity criteria or is less than 8 characters.")
	}
}

