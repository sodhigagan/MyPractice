package main

import "fmt"

func main() {
	var text string
	fmt.Print("Enter text: ")
	fmt.Scan(&text)
	fmt.Println("text:", text)
	/*for i := 0; text[i] != ' '; i++ {
		fmt.Println(string(text[i]))
	}*/
}
