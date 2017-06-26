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
	q := r.URL.Query()
	emailStr := q.Get("email")
	fmt.Println(emailStr)
	if emailStr == "sam" {
		fmt.Println("progress!")
	} else {
		fmt.Println("hmmm...")

	}

}

func ValidateEmail2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Currently Validating email\n"))
	q := mux.Vars(r)
	emailStr := q["email"]
	fmt.Println(emailStr)
	if emailStr == "sam" {
		fmt.Println("progress!")
	} else {
		fmt.Println("hmmm...")

	}

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", YourHandler)
	r.HandleFunc("/validate", ValidateEmail)
	r.HandleFunc("/validate/{email}", ValidateEmail2)
	log.Fatal(http.ListenAndServe(":9002", r))
}
