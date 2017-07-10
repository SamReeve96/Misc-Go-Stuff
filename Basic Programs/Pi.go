package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	//Assign variable Pi the value of Pi (Float)
	pi := math.Pi
	fmt.Println("Pi Equals", pi)
	//Intialize rand generator
	rand.Seed(12341234123)
	fmt.Println("random number", rand.Intn(10))
}
