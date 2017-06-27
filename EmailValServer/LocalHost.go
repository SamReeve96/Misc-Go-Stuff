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

func ObtainEmail(w http.ResponseWriter, r *http.Request) {
	q := mux.Vars(r)
	emailStr := q["email"]
	fmt.Println(emailStr)
	if len(emailStr) > 0 {
		valid := ValidateEmail(emailStr)
		if valid == true {
			w.Write([]byte("Your email is valid!\n"))
		} else {
			w.Write([]byte("Your email is not valid\n"))
		}
	} else {
		w.Write([]byte("String passed is empty, try again\n URL should look like: `/validate?email=example@eg.com`"))
	}

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", YourHandler)
	r.HandleFunc("/validate/{email}", ObtainEmail)
	log.Fatal(http.ListenAndServe(":9003", r))
}
