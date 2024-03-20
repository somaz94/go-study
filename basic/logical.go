package main

import "fmt"

func logical() {
    A := true
    B := false
    C := true
    fmt.Println("A && B:", A && B)
    fmt.Println("A || !(C && B):", A || !(C && B))
}
