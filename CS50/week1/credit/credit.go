package main

import "fmt"

func main() {
	var number, ccnum, ccnum1, sum, sum1, sum2, seclastdigit, digits, digits1 int
	fmt.Print("Credit Card Number: ")
	fmt.Scan(&number)

	digits = 0
	ccnum = number
	ccnum1 = number
	sum1 = 0
	sum2 = 0

	// Calculating the digits here
	for number != 0 {
		number = number / 10
		digits = digits + 1
	}

	// This is Luhn's Algorithm
	for ccnum != 0 {
		sum1 = sum1 + (ccnum % 10)
		ccnum = ccnum / 10
		seclastdigit = (ccnum % 10) * 2
		for seclastdigit != 0 {
			sum2 = sum2 + (seclastdigit % 10)
			seclastdigit = seclastdigit / 10
		}
		ccnum = ccnum / 10
	}
	sum = sum1 + sum2

	if sum%10 == 0 {
		for digits1 = digits - 2; digits1 > 0; digits1-- {
			ccnum1 = ccnum1 / 10
		}
		if ccnum1 == 37 && digits == 15 {
			fmt.Println("Amex")
		} else if (digits == 13 || digits == 16) && ccnum1/10 == 4 {
			fmt.Println("Visa")
		} else if (digits == 16) && (ccnum1 == 51 || ccnum1 == 52 || ccnum1 == 53 || ccnum1 == 54 || ccnum1 == 55) {
			fmt.Println("Mastercard")
		}
	} else {
		fmt.Println("Invalid")
	}
}
