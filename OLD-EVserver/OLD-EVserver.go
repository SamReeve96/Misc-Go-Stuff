/*
	This file demonstrates two ways of obtaining a string from the URL either by
	a url Query and get (ValidateEmail),
	or using gorilla mux (ValidateEmail2) as seen below.
*/

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func ValidateEmail(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Currently Validating email\n"))
	q := r.URL.Query()
	emailStr := q.Get("email")
	fmt.Println(emailStr)
	if emailStr == "Test" {
		fmt.Println("progress!")
	} else {
		fmt.Println("error")
	}
}

func ValidateEmail2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Currently Validating email\n"))
	q := mux.Vars(r)
	emailStr := q["email"]
	fmt.Println(emailStr)
	if emailStr == "Test" {
		fmt.Println("progress!")
	} else {
		fmt.Println("error")

	}

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/validate", ValidateEmail)
	r.HandleFunc("/validate/{email}", ValidateEmail2)
	log.Fatal(http.ListenAndServe(":9002", r))
}
