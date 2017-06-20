package main

import (
	"fmt"
	"regexp"
)

func ValidateEmail(emailStr string) bool {
	IsValid := false
	regexStr := `[a-z0-9]+@[a-z0-9]+\.[a-z]+`
	//emailStr := "sreeve96@gmail.com"
	if regexp.MustCompile(regexStr).MatchString(emailStr) == true {
	} else {
		IsValid = false
	}
	return IsValid
}

func main() {
	fmt.Println(ValidateEmail("sreeve96@gfkgv897987mail.com"))
	
}
