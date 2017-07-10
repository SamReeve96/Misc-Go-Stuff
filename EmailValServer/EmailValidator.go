package EVLocalHost

import (
	"regexp"
)

//ValidateEmail will check to see if the email passed fits the regex
func ValidateEmail(emailStr string) bool {
	isValid := false
	regexStr2 := ("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	//Change the value of isValid based on the result of passing the email string through the Regex
	isValid = regexp.MustCompile(regexStr2).MatchString(emailStr)
	return isValid
}
