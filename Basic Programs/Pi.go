package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	Pi := math.Pi
	fmt.Println("Pi Equals", Pi)
	rand.Seed(12341234123)
	fmt.Println("random number", rand.Intn(10))
}
