package EmailValidator

import (
	"strconv"
	"testing"
)

func TestEmailVal(t *testing.T) {
	//Create a slice containing serveral test cases
	cases := []struct {
		email     string
		expResult bool
	}{
		{"sreeve96@gmail.com", true},
		{"accountxboxsams@hotmail.com", true},
		{"Sam.James@Reeve.com", true},
		{"_____@gmail.ac.uk", true},
		{"Sam-James@Reeve.com", true},
		{"lorum ipsum", false},
		{"sr@eve96@$£%$^.com", false},
		{"email@e*ample.fail", false},
		{"Sam@Reeve..com", false},
	}

	//For in range in go can use the index (what _ is)or value of the object in the data store
	//being iterated through (what C is)
	for _, c := range cases {
		//Pass the value to the validateEmail function in the main .go file
		result := ValidateEmail(c.email)
		//If the result is not what was expected, the test has failed, Log this and tell the user
		if result != c.expResult {
			t.Log("Should be: " + strconv.FormatBool(c.expResult) + " got: " + strconv.FormatBool(result))
			t.Fail()
		}
	}
}
