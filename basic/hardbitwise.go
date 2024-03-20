package main

import "fmt"

func hardbitwise() {
    // Define two integers
    a := 9  // binary: 1001
    b := 5  // binary: 0101

    // Bitwise AND
    andResult := a & b  // This will be 1 (binary: 0001), since only the last bit is 1 in both numbers
    fmt.Printf("%d AND %d = %d\n", a, b, andResult)

    // Bitwise OR
    orResult := a | b  // This will be 13 (binary: 1101), since bits from either number are 1
    fmt.Printf("%d OR %d = %d\n", a, b, orResult)

    // Left shift (a << 2)
    leftShift := a << 2  // Shifts the bits of a two places to the left (1001 becomes 100100), which is 36 in decimal
    fmt.Printf("%d left shift by 2 = %d\n", a, leftShift)

    // Right shift (a >> 2)
    rightShift := a >> 2  // Shifts the bits of a two places to the right (1001 becomes 10), which is 2 in decimal
    fmt.Printf("%d right shift by 2 = %d\n", a, rightShift)
}
