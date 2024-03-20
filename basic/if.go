package main

import "fmt"

func iffunction() {
	// Let's check this number
	number := -5

	// Using if, else if, and else to check the number's sign
	if number > 0 {
		fmt.Println(number, "is positive")
	} else if number < 0 {
		fmt.Println(number, "is negative")
	} else {
		fmt.Println(number, "is zero")
	}
}
