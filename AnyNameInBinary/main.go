package main

import (
	"fmt"
)

func main() {
	var name string
	var length, i int
	fmt.Println("Enter your name:")
	fmt.Scan(&name)
	length = len(name)
	fmt.Print("Nice to meet you ")

	for i = 0; i < length; i++ {
		fmt.Printf("%b ", name[i])
	}
}
