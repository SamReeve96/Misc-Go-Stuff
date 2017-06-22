package EmailValidator

import (
	"strconv"
	"testing"
)

func TestEmailVal(t *testing.T) {
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

	for _, c := range cases {
		result := ValidateEmail(c.email)
		if result != c.expResult {
			t.Log("Should be: " + strconv.FormatBool(c.expResult) + " got: " + strconv.FormatBool(result))
			t.Fail()
		}
	}
}
