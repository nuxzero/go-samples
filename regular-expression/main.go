package main

import (
	"fmt"
	"regexp"
)

func main() {
	isEnglishName := validateEnglishName("John Doe")
	fmt.Printf("isEnglishName: %v\n", isEnglishName)

	isPhoneNumber := validatePhoneNumber("0123456789")
	fmt.Printf("isPhoneNumber: %v\n", isPhoneNumber)

	isEmail := validateEmail("johndoe@email.com")
	fmt.Printf("isEmail: %v\n", isEmail)
}

func validateEnglishName(name string) bool {
	re := regexp.MustCompile(`^[a-zA-Z\s]+$`)
	return re.MatchString(name)
}

func validatePhoneNumber(number string) bool {
	re := regexp.MustCompile(`^0[0-9]{8,9}$`)
	return re.MatchString(number)
}

func validateEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func extractStreetName(address string) string {
	re := regexp.MustCompile(`, ([a-zA-Z0-9\s]+) (Rd|Rd.|Road),`)
	matches := re.FindStringSubmatch(address)
	if len(matches) < 2 {
		return ""
	}
	return matches[1]
}
