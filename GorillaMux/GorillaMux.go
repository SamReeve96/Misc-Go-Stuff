package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!\n"))
}

func ValidateEmail(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Currently Validating email\n"))
	vars := mux.Vars(r)
	emailStr := vars["emails"]
	if emailStr == "sam" {
		fmt.Println("progress!")
	} else {
		fmt.Println("hmmm...")
		fmt.Print(emailStr)
	}

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", YourHandler)
	r.HandleFunc("/validate", ValidateEmail)
	log.Fatal(http.ListenAndServe(":9002", r))
}
