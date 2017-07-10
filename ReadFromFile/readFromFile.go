//This program reads in a whole .txt file.
//It then stores the file line by line in a slice

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fileReader()
}

func fileReader() []string {
	fullFile, err := os.Open("Emails.txt")
	errCheck(err)
	var lines []string
	scanner := bufio.NewScanner(fullFile)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		fmt.Println(lines)
	}
	errCheck(err)
	return lines
}

func errCheck(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
