package main

import "fmt"

func arrays() {
	// 1. Declaration and Initialization of a Single-Dimensional Array
	var singleArray [3]int = [3]int{1, 2, 3}
	fmt.Println("Single-Dimensional Array:", singleArray)

	// Using shorthand declaration
	shorthandArray := [...]int{4, 5, 6}
	fmt.Println("Shorthand Array:", shorthandArray)

	// 2. Accessing Array Elements
	fmt.Println("Second element of Single-Dimensional Array:", singleArray[1])

	// 3. Declaration and Initialization of a Multi-Dimensional Array
	var multiArray [2][3]int = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("Multi-Dimensional Array:", multiArray)

	var threeDArray [3][3][3]int = [3][3][3]int{
		{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
		{
			{10, 11, 12},
			{13, 14, 15},
			{16, 17, 18},
		},
		{
			{19, 20, 21},
			{22, 23, 24},
			{25, 26, 27},
		},
	}
	fmt.Println("Three-Multi-Dimensional Array:", threeDArray)

	// 4. Accessing Multi-Dimensional Array Elements
	fmt.Println("Access element [1][2] of Multi-Dimensional Array:", multiArray[1][2])

	// Demonstrating automatic array size determination
	autoSizeArray := [...]int{7, 8, 9, 10}
	fmt.Println("Array with automatically determined size:", autoSizeArray)
	fmt.Printf("Type of autoSizeArray: %T\n", autoSizeArray) // Displays the type of the array, showing its size

	// 5. Modifying Array Elements
	singleArray[0] = 100
	fmt.Println("Modified Single-Dimensional Array:", singleArray)

	// 6. Looping Through Array Elements
	fmt.Println("Iterating over singleArray:")
	for index, value := range singleArray {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// 7. Exploring the Effect of Array Size on Type
	fmt.Println("Exploring Array Types:")
	var arrayTypeOne [3]int
	var arrayTypeTwo [4]int
	fmt.Printf("Type of arrayTypeOne: %T\n", arrayTypeOne)
	fmt.Printf("Type of arrayTypeTwo: %T\n", arrayTypeTwo)

	// Note: arrayTypeOne and arrayTypeTwo are considered of different types because of their different sizes.
	// Note: fmt.Println은 자동 개행 fmt.printf는 개행을 해줘야한다.
}
