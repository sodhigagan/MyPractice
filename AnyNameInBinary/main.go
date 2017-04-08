package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Enter your name: ")
	absorber := bufio.NewScanner(os.Stdin)
	absorber.Scan()
	name := absorber.Text()
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
