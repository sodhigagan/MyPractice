package main

import "fmt"

var change, cents float32

func main() {
	fmt.Print("O hai!")
	fmt.Print("How much change is owed: ")
	fmt.Scan(&change)

	cents = (change * 100)
	fmt.Println("Cents are ", cents)
	fmt.Printf("Cents are %.2f", cents)
}
