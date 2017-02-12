package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter your name:")
	letter := bufio.NewScanner(os.Stdin)
	letter.Scan()
	a := letter.Text()
	fmt.Println("So your name is: ", a) //, letter.Text())
	//fmt.Printf("%b", a)
}
