package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Enter your name: ")
	absorber := bufio.NewScanner(os.Stdin) // input
	absorber.Scan()                        // I am not familiar with this usage, however the input is being scanned??
	name := absorber.Text()                // input is being assigned to 'name', I don't know why I can't just use 'absorber'
	// var name string I had used this with fmt.Scan, will try it later
	var length, i int
	//fmt.Scan(&name) since this is not working, i'll try bufio
	length = len(name)
	// fmt.Println("Length of", name, "is ", length) testing
	fmt.Print("Nice to meet you ")

	for i = 0; i < length; i++ {
		fmt.Printf("%b ", name[i])
	}
}
