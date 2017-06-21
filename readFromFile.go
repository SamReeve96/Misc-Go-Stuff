//Thanks to openwonk on Stackoverflow.com

package main

import (
	"bufio"
	"fmt"
	//"io"
	//"io/ioutil"
	"os"
	//"strings"
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
