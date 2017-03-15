package main

import "fmt"

func main() {
	var i, sum int
	for i = 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum = sum + i
		}
	}
	fmt.Println(sum)
}
