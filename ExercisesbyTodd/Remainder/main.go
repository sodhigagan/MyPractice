package main

import "fmt"

func main() {
	var small, large, remainder int
	fmt.Print("Enter the smaller number: ") //there ain't any sanity check here
	fmt.Scan(&small)
	fmt.Print("Enter the larger number: ") //there ain't any sanity check here
	fmt.Scan(&large)
	remainder = large % small
	fmt.Println("After dividing ", large, "by", small, "The remainder is ", remainder)
}
