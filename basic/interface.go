package main

import (
	"fmt"
	"math"
)

// Shape interface with area() and perimeter() methods
type Shape interface {
	area() float64
	perimeter() float64
}

// Rect struct
type Rects struct {
	width, height float64
}

// Circle struct
type Circle struct {
	radius float64
}

// Implementation of Shape interface for Rect
func (r Rects) area() float64 {
	return r.width * r.height
}

func (r Rects) perimeter() float64 {
	return 2 * (r.width + r.height)
}

// Implementation of Shape interface for Circle
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// Function accepting Shape interfaces as parameters
func showArea(shapes ...Shape) {
	for _, shape := range shapes {
		fmt.Println("Area:", shape.area())
	}
}

// interfaces demonstrates working with interfaces and empty interfaces
func interfaces() {
	rects := Rects{width: 10, height: 20}
	circle := Circle{radius: 10}

	showArea(rects, circle)

	// Working with empty interfaces
	var x interface{}
	x = "Tom"
	printIt(x)

	// Type Assertion
	var a interface{} = 1
	j := a.(int)
	fmt.Println("Type asserted value:", j)
}

// Function to work with empty interface
func printIt(v interface{}) {
	fmt.Println("Value in empty interface:", v)
}
