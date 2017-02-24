package main

import "fmt"

var a string

func main() {
	fmt.Scan(&a)
	fmt.Println(a, " is stored at ", &a)

	var b = &a
	fmt.Println("b's address is ", b)
	*b = "Hello, world"
	fmt.Println("However the address has been updated with ", a)
}
