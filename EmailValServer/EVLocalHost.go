package EVLocalHost

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Basic function to check that the local host is working by writing hello on the page
func ashensEasterEgg(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!\n"))
}

//Will obtain the email from the URL and pass it to the function to validate it, then display the result
func valEmail(w http.ResponseWriter, r *http.Request) {
	//Get the URl
	q := r.URL.Query()
	//Store the part of the url that corresponds to the Get
	emailStr := q.Get("email")
	if len(emailStr) > 0 {
		valid := ValidateEmail(emailStr)
		if valid == true {
			w.Write([]byte("Your email is valid"))
		} else {
			w.Write([]byte("Your email is not valid"))
		}
	} else {
		w.Write([]byte("String passed is empty, try again"))
	}

}

func main() {
	//create a new router
	r := mux.NewRouter()
	//Comment out the passEmail if it's not working as expected and see if the hello function is working
	r.HandleFunc("/hello", ashensEasterEgg)
	r.HandleFunc("/validate/{email}", valEmail)
	//Serve the page and listen for the paths that are passed
	log.Fatal(http.ListenAndServe(":9003", r))
}
