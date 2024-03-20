package main

import "fmt"

func hardiffunction() {
    // Let's check these numbers
    numbers := []int{-5, 0, 4}

    for _, number := range numbers {
        // Check if the number is positive, negative, or zero
        if number > 0 {
            fmt.Print(number, " is positive")
        } else if number < 0 {
            fmt.Print(number, " is negative")
        } else {
            fmt.Println(number, "is zero")
            continue // Skip the even/odd check for zero
        }

        // Check if the number is even or odd
        // This part is skipped if the number is zero
        if number%2 == 0 {
            fmt.Println(" and even")
        } else {
            fmt.Println(" and odd")
        }
    }
}
