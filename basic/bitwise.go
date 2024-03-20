package main

import "fmt"

func bitwise() {
	a := 2 // binary: 10
	b := 3 // binary: 11
	c := (a & b) << 5
	fmt.Println("Result of (a & b) << 5:", c)
}
