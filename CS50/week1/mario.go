package main

import "fmt"

var h int

func main() {
Height:
	fmt.Print("Height: ")
	fmt.Scan(&h)

	if h < 0 {
		fmt.Println("too less, try again")
		goto Height
	} else if h > 20 {
		fmt.Println("too much, try again")
		goto Height
	} else {
		fmt.Println("perfect")
	}
}
