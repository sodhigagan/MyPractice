package main

import "fmt"

var change float32

func main() {
	fmt.Print("O hai! ")
	for {
		fmt.Print("How much change is owed: ")
		fmt.Scan(&change)
		if change <= 0 {
			fmt.Print("That is not correct value, try again! ")
			continue
		} else if change > 0 {
			break
		}
	}
	var cents, loonie, quarter, dime, nickel, penny, coins int

	change = change * 100
	cents = int(change) //this just drops everything after the decimal
	//fmt.Println("First Value: ", cents)

	if cents > 100 {
		loonie = cents / 100
		cents = cents % 100
		//fmt.Println("Loonification: ", cents)
	}
	if cents > 25 && cents < 100 {
		quarter = cents / 25
		cents = cents % 25
		//fmt.Println("Quarterfication: ", cents)
	}
	if cents > 10 && cents < 25 {
		dime = cents / 10
		cents = cents % 10
		//fmt.Println("Dimification: ", cents)
	}
	if cents > 5 && cents < 10 {
		nickel = cents / 5
		cents = cents % 5
		//fmt.Println("Centification: ", cents)
	}
	if cents > 1 && cents < 5 {
		penny = cents / 1
		//cents = cents % 1 because I am not using this calculaion
		//fmt.Println("Pennification: ", cents)
	}

	coins = loonie + quarter + dime + nickel + penny
	fmt.Println("Loonies:", loonie)
	fmt.Println("Quarters:", quarter)
	fmt.Println("Dimes:", dime)
	fmt.Println("Nickels:", nickel)
	fmt.Println("Pennies:", penny)
	fmt.Println("Total coins:", coins)
}
