package EmailValidator

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
	if IsValid == true {
		fmt.Println(emailStr, "Address is valid")
	} else {
		fmt.Println(emailStr, "Address is not valid")
	}
	return IsValid
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter an address")
	emailStr, _ := reader.ReadString('\n')
	ValidateEmail(emailStr)
}
