package EVLocalHost

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPassEmail(t *testing.T) {

	cases := []struct {
		url       string
		expResult bool
	}{
		{"/Validate?email=sreeve96@gmail.com", true},
		{"/Validate?email=accountxboxsams@hotmail.com", true},
		{"/Validate?email=Sam.James@Reeve.com", true},
		{"/Validate?email=_____@gmail.ac.uk", true},
		{"/Validate?email=Sam-James@Reeve.com", true},
		{"/Validate?email=email@e*ample.fail", false},
		{"/Validate?email=Sam@Reeve..com", false},

		{"/Validate?email Sam@reeve.com", false},
		{"/email=sr@eve96.com", false},
	}

	for _, c := range cases {
		req, err := http.NewRequest("GET", c.url, nil)
		if err != nil {
			t.Fatal(err)
		}

		//This creates a rr (responce recorder)
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(PassEmail)

		handler.ServeHTTP(rr, req)

		//now call the serveHTTP directly and pass the request and responce recorder
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Handler returned the wrong status code, \n Got: %v  \n Want: %v", status, http.StatusOK)
		}

		//Use result boolean as this is easier to test when using multiple cases. (Intialise result bool)
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
