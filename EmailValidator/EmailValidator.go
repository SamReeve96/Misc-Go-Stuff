package EmailValidator

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

//ValidateEmail will check to see if the email passed fits the regex
func ValidateEmail(emailStr string) bool {
	isValid := false
	regexStr := `[a-z0-9]+@[a-z0-9]+\.[a-z]+`
	isValid = regexp.MustCompile(regexStr).MatchString(emailStr)
	if isValid == true {
		fmt.Println(emailStr, "Address is valid")
	} else {
		fmt.Println(emailStr, "Address is not valid")
	}
	return isValid
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter an address")
	emailStr, _ := reader.ReadString('\n')
	ValidateEmail(emailStr)
}
