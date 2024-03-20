package main

import "fmt"

func assignment() {
    a := 100
    a *= 10
    a >>= 2
    a |= 1
    fmt.Println("Final value of a:", a)
}
