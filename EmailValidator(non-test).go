package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

//ValidateEmail will check to see if the email passed fits the regex
func ValidateEmail(emailStr string) {
	IsValid := false
	regexStr := `[a-z0-9]+@[a-z0-9]+\.[a-z]+`
	if regexp.MustCompile(regexStr).MatchString(emailStr) == true {
		IsValid = true
	} else {
		IsValid = false
	}
	if IsValid == true {
		fmt.Println(emailStr, "is valid")
	} else {
		fmt.Println(emailStr, "is not valid")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter an address")
	emailStr, _ := reader.ReadString('\n')
	ValidateEmail(emailStr)
}
