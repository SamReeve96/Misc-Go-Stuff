package EVLocalHost

import (
	"regexp"
)

//ValidateEmail will check to see if the email passed fits the regex
func ValidateEmail(emailStr string) bool {
	isValid := false
	//regexStr := `[a-z0-9]+@[a-z0-9]+\.[a-z]+`
	regexStr2 := ("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	isValid = regexp.MustCompile(regexStr2).MatchString(emailStr)

	/*
		Was useful for terminal based inventory but not needed if web based

		if isValid == true {
			fmt.Println(emailStr, "Address is valid")
		} else {
			fmt.Println(emailStr, "Address is not valid")
		}
	*/
	return isValid
}
