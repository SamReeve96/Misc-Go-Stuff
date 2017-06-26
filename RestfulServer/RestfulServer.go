package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Payload struct {
	Stuff Data
}

type Data struct {
	Fruit Fruits
	Veg   Vegetables
}

type Fruits map[string]int
type Vegetables map[string]int

func serverRest(w http.ResponseWriter, r *http.Request) {
	responce, err := getJSONResponce()
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, string(responce))
}

func main() {
	http.HandleFunc("/", serverRest)
	http.ListenAndServe("localhost:9001", nil)
	fmt.Println("local host should be running on localhost:9001")
}

func getJSONResponce() ([]byte, error) {
	fruits := make(map[string]int)
	fruits["Apples"] = 25
	fruits["Oranges"] = 11

	Vegetables := make(map[string]int)
	Vegetables["Carrots"] = 21
	Vegetables["Peppers"] = 0

	d := Data{fruits, Vegetables}
	p := Payload{d}

	return json.MarshalIndent(p, "", "  ")
}
