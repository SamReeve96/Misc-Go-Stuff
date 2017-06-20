package main

import (
    "fmt"
    "math"
	"math/rand"
)

func main() {
	fmt.Println("Pi Equals", math.Pi)
	rand.Seed(12341234123)
	fmt.Println("random number", rand.Intn(10))
}