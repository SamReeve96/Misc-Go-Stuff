package main

import (
	"fmt"
	"regexp"
)

//ValidateEmail will check to see if the email passed fits the regex
func ValidateEmail(emailStr string) bool {
	IsValid := false
	regexStr := `[a-z0-9]+@[a-z0-9]+\.[a-z]+`
	//altemailStr := "sreeve96@gmail.com"
	if regexp.MustCompile(regexStr).MatchString(emailStr) == true {
		IsValid = true
	} else {
		IsValid = false
	}
	return IsValid
}

func main() {
	fmt.Println(ValidateEmail("sreeve96@gmail.com"))

}
