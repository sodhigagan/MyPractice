package main

import "fmt"

func main() {
	var i int
	for i = 1; i <= 100; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			(fmt.Println("FizzBuzz"))
		case i%3 == 0:
			(fmt.Println("Fizz"))
		case i%5 == 0:
			(fmt.Println("Buzz"))
			/*case i%3 != 0 || i%5 != 0:*/
		default:
			(fmt.Println(i))
		}
	}
}
