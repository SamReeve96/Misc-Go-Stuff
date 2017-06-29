package EVLocalHost

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!\n"))
}

func PassEmail(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
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
	r := mux.NewRouter()
	r.HandleFunc("/hello", YourHandler)
	r.HandleFunc("/validate/{email}", PassEmail)
	log.Fatal(http.ListenAndServe(":9003", r))
}
