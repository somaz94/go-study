package main

import "fmt"

func switchs() {
	var i interface{} = "hello"

	// No expression after switch
	// This uses the fact that the zero value for an integer is 0
	var num = 3
	switch {
	case num < 0:
		fmt.Println(num, "is negative")
	case num == 0:
		fmt.Println(num, "is zero")
	default:
		fmt.Println(num, "is positive")
	}

	// Using expressions in case statements
	switch age := 25; {
	case age < 18:
		fmt.Println("Minor")
	case age >= 18 && age < 65:
		fmt.Println("Adult")
	default:
		fmt.Println("Senior")
	}

	// Handling multiple values in a single case and using fallthrough
	month := 5
	switch month {
	case 12, 1, 2:
		fmt.Println("Winter")
		fallthrough
	case 3, 4, 5:
		fmt.Println("Spring")
		if month == 5 {
			fmt.Println("It's May!")
		}
	case 6, 7, 8:
		fmt.Println("Summer")
	case 9, 10, 11:
		fmt.Println("Fall")
	default:
		fmt.Println("Invalid month")
	}

	// Type switch
	switch v := i.(type) {
	case int:
		fmt.Println("Integer:", v)
	case bool:
		fmt.Println("Boolean:", v)
	case string:
		fmt.Println("String:", v)
	default:
		fmt.Printf("Unknown type %T\n", v)
	}
}
