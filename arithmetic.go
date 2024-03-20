package main

import "fmt"

func arithmetic() {
    a := 10
    b := 5
    c := (a + b) / 5
    a++ // increment a by 1
    fmt.Println("Result of (a + b) / 5:", c)
    fmt.Println("a after increment:", a)
}

