package EmailValidator

import (
	//"strings"
	"testing"
)

func TestEmailVal(t *testing.T) {
	result := ValidateEmail("sreeve96@gmail.com")
	if result != true {
		t.Log("Should be true, got:", result)
		t.Fail()
	}
}
