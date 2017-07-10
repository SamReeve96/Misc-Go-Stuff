package EmailValidator

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

//ValidateEmail will check to see if the email passed is valid to the regex
func ValidateEmail(emailStr string) bool {
	//Intallize boolean
	isValid := false
	regexStr := ("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	//Update the value of isValid based on whether the string passed satisfies the regex
	isValid = regexp.MustCompile(regexStr).MatchString(emailStr)
	if isValid == true {
		fmt.Println(emailStr, "Address is valid")
	} else {
		fmt.Println(emailStr, "Address is not valid")
	}
	return isValid
}

func main() {
	//Create a reader to get the user's input
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter an address")
	//Read the users input and store it in the emailStr variable
	//(_ is an unassigned variable that here is used to store the err string)
	emailStr, _ := reader.ReadString('\n')
	//Pass the string to the validateEmail method
	ValidateEmail(emailStr)
}
