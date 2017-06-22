package EmailValidator

import (
	//"strings"
	"fmt"
	"testing"
)

func TestEmailVal(t *testing.T) {
	result, err := ValidateEmail("sreeve96@gmail.com")
	if result != true {
		t.Log("Should be true, got:", result)
	} else {
		fmt.Println(err)
	}
}
