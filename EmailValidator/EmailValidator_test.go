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
		{"lorum ipsum", false},
		{"accountxboxsams@hotmail.com", true},
		{"sr@eve96@$£%$^.com", false},
	}

	for _, c := range cases {
		result := ValidateEmail(c.email)
		if result != c.expResult {
			t.Log("Should be: " + strconv.FormatBool(c.expResult) + " got: " + strconv.FormatBool(result))
			t.Fail()
		}
	}
}
