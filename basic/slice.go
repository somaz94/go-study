package main

import "fmt"

func slices() {
	// Initialize slices
	sliceA := []int{0, 1}
	sliceB := []int{4, 5, 6}

	// Adding elements to a slice
	sliceA = append(sliceA, 2)       // Adds a single element
	sliceA = append(sliceA, 3, 4, 5) // Adds multiple elements

	// Merging slices
	sliceA = append(sliceA, sliceB...) // Merges sliceB into sliceA

	fmt.Println("After adding and merging:", sliceA)

	// Copying slices
	source := []int{7, 8, 9}
	target := make([]int, len(source), cap(source)*2)
	copy(target, source)
	fmt.Println("After copying:", target)

	// Demonstrating slice capacity behavior
	dynamicSlice := make([]int, 0, 3) // slice with len=0, cap=3
	for i := 1; i <= 15; i++ {
		dynamicSlice = append(dynamicSlice, i)
		// Outputting length and capacity to observe changes
		if i == 1 || i == 3 || i == 6 || i == 12 || i == 15 {
			fmt.Printf("After adding %d: len=%d, cap=%d\n", i, len(dynamicSlice), cap(dynamicSlice))
		}
	}
	fmt.Println("Final slice:", dynamicSlice)
}

// korean
// 초기화 및 추가: 슬라이스를 초기화하고 append()를 사용하여 단일 또는 여러 요소를 추가하는 방법을 보여준다.
// 병합: append()를 사용하지만 줄임표 구문을 사용하여 두 번째 조각의 모든 요소를 ​​추가하여 두 조각을 병합하는 방법을 보여준다.
// 복사: copy() 함수는 한 슬라이스(source)에서 다른 슬라이스(target)로 요소를 복사하는 데 사용되며, 특정 용량으로 대상 슬라이스를 초기화하는 방법을 보여준다.
// 용량 관리: 'append()'가 포함된 루프는 Go가 슬라이스 용량을 자동으로 관리하는 방법을 보여준다. 요소가 현재 용량 이상으로 추가되면 슬라이스의 용량이 동적으로 증가하며, 새 요소를 효율적으로 수용하기 위해 종종 두 배로 늘어난다.
// 관찰: 특정 지점의 길이와 용량을 인쇄함으로써 이 코드는 슬라이스가 커짐에 따라 용량이 언제 어떻게 변하는지 관찰하는 데 도움이 된다.

// func append(slice []Type, elems ...Type) []Type

// english
// Initialization and Addition: It shows how to initialize slices, add single or multiple elements using append().
// Merging: It demonstrates how to merge two slices together, also using append() but with the ellipsis syntax to add all elements of the second slice.
// Copying: The copy() function is used to copy the elements from one slice (source) to another (target), showcasing how to initialize the target slice with a specific capacity.
// Capacity Management: The loop with append() illustrates how Go automatically manages slice capacity. It increases the capacity of the slice dynamically as elements are added beyond its current capacity, often doubling it to accommodate new elements efficiently.
// Observation: By printing out the length and capacity at certain points, this code helps us observe how and when the capacity changes as the slice grows.
