package main

import "fmt"

func pointer() {
    k := 10
    p := &k // pointer to k
    fmt.Println("Address of k:", p)
    fmt.Println("Value of k through pointer:", *p)
}
