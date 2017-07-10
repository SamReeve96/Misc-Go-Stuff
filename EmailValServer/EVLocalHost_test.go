package EVLocalHost

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPassEmail(t *testing.T) {
	//Create a slice containing the test strings and its corresponding expected result
	cases := []struct {
		url       string
		expResult bool
	}{
		//This block of tests should pass through to the validator
		{"/Validate?email=sreeve96@gmail.com", true},
		{"/Validate?email=accountxboxsams@hotmail.com", true},
		{"/Validate?email=Sam.James@Reeve.com", true},
		{"/Validate?email=_____@gmail.ac.uk", true},
		{"/Validate?email=Sam-James@Reeve.com", true},
		{"/Validate?email=email@e*ample.fail", false},
		{"/Validate?email=Sam@Reeve..com", false},
		//These tests shouldn't pass to the validator as their URL's are the error
		{"/Validate?email Sam@reeve.com", false},
		{"/email=sr@eve96.com", false},
	}

	for _, c := range cases {
		req, err := http.NewRequest("GET", c.url, nil)
		if err != nil {
			t.Fatal(err)
		}

		//This creates a rr (response recorder)
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(valEmail)
		//now call the serveHTTP directly and pass the request and response recorder
		handler.ServeHTTP(rr, req)

		//check to see if this was successful, if not show an error
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Handler returned the wrong status code, \n Got: %v  \n Want: %v", status, http.StatusOK)
		}

		//capture the body of the webpage, determine the value of a boolean based on the body string
		result := false
		if rr.Body.String() == "Your email is valid" {
			result = true
			//If the error is with the url and returned an empty sting, tell the user
		} else if rr.Body.String() == "String passed is empty, try again" {
			fmt.Printf("The URL: %v is incorrect \n", c.url)
		}

		//Now check the body is what is expected
		expected := c.expResult
		if result != expected {
			t.Errorf("Handler returned unexpected body,\n Browser Said: %v \n URL passed: %v \n Result was: %v \n Expected: %v", rr.Body.String(), c.url, result, c.expResult)
		}
	}
}
