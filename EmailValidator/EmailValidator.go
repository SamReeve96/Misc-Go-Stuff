package EmailValidator

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
)

//ValidateEmail will check to see if the email passed fits the regex
func ValidateEmail(emailStr string) (bool, error) {
	isValid := false
	regexStr := `[a-z0-9]+@[a-z0-9]+\.[a-z]+`
	if regexp.MustCompile(regexStr).MatchString(emailStr) == true {
		isValid = true
	} else {
		return false, errors.New("error occured")
	}
	if isValid == true {
		fmt.Println(emailStr, "Address is valid")
	} else {
		fmt.Println(emailStr, "Address is not valid")
	}
	return false, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter an address")
	emailStr, _ := reader.ReadString('\n')
	ValidateEmail(emailStr)
}
