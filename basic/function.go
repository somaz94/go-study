package main

func functions() {
	// Function usage
	message := "Hello, World!"
	say(message)

	// Pass by reference
	msg := "Hello"
	modifyMessage(&msg)
	println("Modified Message:", msg)

	// Variadic function usage
	numbers := []int{1, 2, 3, 4, 5}
	printNumbers(numbers...)

	// Function return value
	count, total := calculateSum(1, 2, 3, 4, 5)
	println("Count:", count, "Total:", total)
}

// A simple function that prints a message
func say(message string) {
	println(message)
}

// Modifies the string value pointed to by the pointer
func modifyMessage(msg *string) {
	*msg = "Changed"
}

// Prints each number from a variadic slice of ints
func printNumbers(nums ...int) {
	for _, num := range nums {
		println(num)
	}
}

// Returns the count and sum of a variadic slice of ints
func calculateSum(nums ...int) (count int, total int) {
	for _, n := range nums {
		total += n
	}
	count = len(nums)
	return // Named return parameters are implicitly returned
}

// korean
// 이 예에서는 간단한 함수 say를 사용하여 메시지를 인쇄하는 것으로 시작한다.
// 다음 원래 문자열을 수정하기 위해 함수에 문자열 포인터를 전달하는 modifyMessage를 사용하여 참조로 전달하는 방법을 보여준다.
// 'printNumbers' 함수는 임의 개수의 정수를 허용하는 가변 함수를 보여준다.
// 마지막으로 'calculateSum'은 여러 값, 특히 일련의 숫자의 개수와 총합을 반환하는 함수를 보여준다.
// 명확성과 편의성을 위해 calculateSum에서 명명된 반환 매개변수를 어떻게 사용하는지 주목한다.

// english
// In this example, we start by using a simple function say to print a message.
// We then demonstrate passing by reference with modifyMessage, where we pass a string pointer to the function to modify the original string.
// The printNumbers function showcases a variadic function that accepts an arbitrary number of integers.
// Finally, calculateSum demonstrates a function that returns multiple values, specifically the count and total sum of a series of numbers.
// Notice how we make use of named return parameters in calculateSum for clarity and convenience.
