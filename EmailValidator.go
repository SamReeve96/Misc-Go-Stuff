package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

//ValidateEmail will check to see if the email passed fits the regex
func ValidateEmail(emailStr string) bool {
	IsValid := false
	regexStr := `[a-z0-9]+@[a-z0-9]+\.[a-z]+`
	if regexp.MustCompile(regexStr).MatchString(emailStr) == true {
		IsValid = true
	} else {
		IsValid = false
	}
	return IsValid
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter an address")
	emailStr, _ := reader.ReadString('\n')
	Result := ValidateEmail(emailStr)
	if Result == true {
		fmt.Println("Address is valid")
	} else {
		fmt.Println("Address is not valid")
	}
}
