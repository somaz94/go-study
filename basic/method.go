package main

import "fmt"

// Rect - struct definition
type Rect struct {
    width, height int
}

// area - Rect's method with a value receiver
func (r Rect) area() int {
    return r.width * r.height
}

// area2 - Rect's method with a pointer receiver
func (r *Rect) area2() int {
    r.width++ // This modifies the Rect's width for the caller
    return r.width * r.height
}

func methods() {
    rect := Rect{10, 20}

    // Call method with value receiver
    areaValueReceiver := rect.area()
    fmt.Println("Using value receiver:", rect.width, areaValueReceiver) // prints: 10 200

    // Call method with pointer receiver
    areaPointerReceiver := rect.area2()
    fmt.Println("Using pointer receiver:", rect.width, areaPointerReceiver) // prints: 11 220

    // Demonstrating the effect of the pointer receiver on the original struct
    fmt.Println("After pointer receiver method call:", rect.width) // prints: 11
}
