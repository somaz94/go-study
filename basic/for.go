package main

import "fmt"

func fors() {
	// Example of using range with a slice
	names := []string{"Alice", "Bob", "Charlie"}
	fmt.Println("Iterating over names with for range:")
	for index, name := range names {
		fmt.Printf("%d: %s\n", index, name)
	}

	// Infinite loop with conditional break and continue
	fmt.Println("\nUsing an infinite loop with break and continue:")
	count := 0
	for {
		count++
		if count%2 == 0 {
			fmt.Printf("%d is even, continue...\n", count)
			if count >= 10 {
				fmt.Println("Count reached 10, breaking...")
				break // Exit the loop when count reaches 10
			}
			continue
		}
		fmt.Printf("%d is odd\n", count)
		if count >= 10 {
			break // Safety break to prevent infinite execution
		}
	}

	// Nested loops with label break
	fmt.Println("\nNested loops with label break:")
	outerLoop := 0
OuterLoop: // Label for the outer loop
	for {
		innerLoop := 0
		for {
			fmt.Printf("OuterLoop: %d, InnerLoop: %d\n", outerLoop, innerLoop)
			if innerLoop == 2 {
				fmt.Println("Breaking out of both loops.")
				break OuterLoop // Break out of the outer loop
			}
			if innerLoop > 3 {
				break // Safety break to prevent infinite inner loop
			}
			innerLoop++
		}
		outerLoop++
		if outerLoop > 3 {
			break // Safety break to prevent infinite outer loop
		}
	}
}

// korean
// for range를 사용하여 슬라이스를 반복하여 색인과 값을 모두 인쇄한다.
// 특정 조건(count >= 10)이 충족되면 continue를 사용하여 짝수를 건너뛰고 break를 사용하여 종료하는 무한 루프이다. 이는 for를 while 루프로 사용하는 방법을 보여준다.
// 내부 루프(innerLoop == 2)에서 특정 조건이 충족될 때 내부 루프와 외부 루프를 모두 벗어나기 위한 레이블(OuterLoop)이 있는 중첩 루프이다. 이것은 레이블과 함께 break를 실제로 사용하는 것이다.

// english
// Iterating over a slice with for range, printing both the index and the value.
// An infinite loop that uses continue to skip even numbers and break to exit when a certain condition (count >= 10) is met. This demonstrates using for as a while loop.
// Nested loops with a label (OuterLoop) to break out of both the inner and outer loop when a certain condition is met in the inner loop (innerLoop == 2). This is a practical usage of break with a label.
