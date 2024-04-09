package main

import "fmt"

// nextValueGenerator returns two closures: one for adding a given value to the counter,
// and another for resetting the counter to a new starting value.
func nextValueGenerator() (func(int) int, func(int)) {
	var counter int = 0
	// Increment the counter by a given value
	addValue := func(val int) int {
		counter += val
		return counter
	}
	// Reset the counter to a new value
	resetCounter := func(val int) {
		counter = val
	}
	return addValue, resetCounter
}

func closures() {
	add, reset := nextValueGenerator()

	fmt.Println(add(3))  // Adds 3; counter is 3
	fmt.Println(add(5))  // Adds 5; counter is 8
	fmt.Println(add(10)) // Adds 10; counter is 18

	// Now let's reset the counter to 5 and start from there
	reset(5)
	fmt.Println(add(2)) // Adds 2; counter is 7
	fmt.Println(add(3)) // Adds 3; counter is 10

	// Creating a new counter
	newAdd, _ := nextValueGenerator()
	fmt.Println(newAdd(1)) // Starts from 0, adds 1; counter is 1
	fmt.Println(newAdd(1)) // Adds 1; counter is 2
}

// korean
// nextValueGenerator 함수: addValue와 resetCounter라는 두 개의 클로저를 반환한다. addValue 클로저는 주어진 값만큼 카운터를 증가시키고 카운터의 새 값을 반환한다. 'resetCounter' 클로저는 카운터를 지정된 새 값으로 재설정한다.
// 클로저 사용: closures 함수는 이러한 클로저를 사용하는 방법을 보여준다. 먼저 'add'를 사용하여 다양한 양만큼 카운터를 증가시킨다. 그런 다음 'reset'을 사용하여 더 많은 값을 추가하기 전에 카운터를 새 값으로 설정한다.
// 새 카운터 생성: nextValueGenerator를 다시 호출하면 클로저가 고유한 상태를 유지하는 방법을 보여주는 새로운 독립 카운터를 생성하는 방법을 보여준다.

// english
// nextValueGenerator Function: Returns two closures - addValue and resetCounter. The addValue closure increments the counter by a given value and returns the new value of the counter. The resetCounter closure resets the counter to a specified new value.
// Using the Closures: The closures function demonstrates using these closures. First, it uses add to increment the counter by various amounts. Then, it uses reset to set the counter to a new value before adding more values.
// Creating a New Counter: Demonstrates that calling nextValueGenerator again creates a new, independent counter, showing how closures maintain their own unique state.
