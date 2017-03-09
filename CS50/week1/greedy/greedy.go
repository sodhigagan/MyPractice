package main

import (
	"fmt"
)

var change float32

func main() {
	fmt.Print("O hai!")
	fmt.Println("How much change is owed: ")
	fmt.Scan(&change)
	fmt.Println("So it is ", change)
}
