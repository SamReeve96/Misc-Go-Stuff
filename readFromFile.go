//Thanks to openwonk on Stackoverflow.com

package main

import (
	"fmt"
	"io/ioutil"
)


func main() {
	fileByte, err := ioutil.ReadFile("Emails.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fileByte) //print as byte format
	fileString := string(fileByte) //Convert from byte to string format
	fmt.Println(fileString) //Print as String
}
