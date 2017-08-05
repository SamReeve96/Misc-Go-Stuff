package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := 1; i <= 100; i++ {
		printString := ""
		fizzResult := check(i, 3)
		buzzResult := check(i, 5)
		if fizzResult {
			printString += "fizz"
		}
		if buzzResult {
			printString += "buzz"
		}
		if !fizzResult && !buzzResult {
			printString += strconv.Itoa(i)
		}
		fmt.Println(printString)
	}
}

func check(number, factor int) (result bool) {
	result = false
	if number%factor == 0 {
		result = true
	}
	return
}
