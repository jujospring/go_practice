package main

import "unicode"

func isValidPassword(password string) bool {
	upperCase := make([]string, 0)
	digits := make([]string, 0)
	chars := make([]string, 0)
	// if len(password) > 12 || len(password) < 5 {
	// 	return false
	// }
	for pos, char := range password {
		if unicode.IsDigit(char) {
			digits = append(digits, string(password[pos]))
		} else if unicode.IsUpper(char) {
			upperCase = append(upperCase, string(password[pos]))
		} else {
			chars = append(chars, string(char))
		}
	}
	numChars := len(digits) + len(upperCase) + len(chars)
	if (numChars >= 5 && numChars <= 12) && (len(digits) != 0 && len(upperCase) != 0) {
		return true
	}
	return false
}
