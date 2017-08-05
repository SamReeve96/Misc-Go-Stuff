package main

import (
	"fmt"
)

func Algorithm() {
	//numbers := []int{}
	fmt.Println("input 10 numbers")
	var num int
	for i := 0; i < 10; i++ {
		fmt.Printf("Debug: i : %d\n", i)
		fmt.Println("Enter next number")
		n, err := fmt.Scanf("%d\n", &num)
		if err != nil {
			fmt.Println(n, err)
		}
		fmt.Println(num)
	}

}

/**
for i := 0; i < sliceSize; i++ {
	var numberToAdd int
	fmt.Printf(" Add number: %d '\n'", i+1)
	n, err := fmt.Scanln("%i\n", &numberToAdd)
	if err != nil {
		fmt.Println(n, err)
	}
	numbers = append(numbers, numberToAdd)
	fmt.Printf("value: %d added", numberToAdd)
} **/
