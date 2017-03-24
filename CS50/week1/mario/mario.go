package main

import "fmt"

func main() {
	var h int
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
		for i := 1; i <= h; i++ {
			for spaces := h - i; spaces != 0; spaces-- {
				fmt.Print(" ")
			}
			for hash1 := h - i; hash1 != h; hash1++ {
				fmt.Print("")
			}
			fmt.Print("  ")
			for hash2 := h - i; hash2 != h; hash2++ {
				fmt.Print("")
			}
			fmt.Print("\n")
		}
	}

}
