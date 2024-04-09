package main

// Define a function prototype for a calculator function
type calculator func(int, int) int

func anonymous_functions() {
	// Anonymous function assigned to a variable
	sum := func(n ...int) int {
		s := 0
		for _, i := range n {
			s += i
		}
		return s
	}
	result := sum(1, 2, 3, 4, 5) // Call the anonymous function
	println("Sum:", result)

	// First-class function: Pass an anonymous function to another function
	add := func(i, j int) int { return i + j }
	r1 := calc(add, 10, 20)
	println("Addition:", r1)

	// Directly pass an anonymous function as a parameter
	r2 := calc(func(x, y int) int { return x - y }, 10, 20)
	println("Subtraction:", r2)

	// Use the calculator prototype to define and pass functions
	multiply := func(x, y int) int { return x * y }
	r3 := calc(multiply, 10, 20)
	println("Multiplication:", r3)
}

// calc uses the calculator prototype for its function parameter
func calc(f calculator, a, b int) int {
	return f(a, b)
}

// korean
// 익명 함수: 가변 개수의 정수 합계를 계산하고 이를 'sum' 변수에 할당하는 익명 함수를 정의한다. 그런 다음 이 함수는 일련의 정수를 인수로 사용하여 호출된다.
// 일급 함수: 'add' 함수는 변수로 정의되어 'calc' 함수에 인수로 전달된다. 이는 함수가 Go에서 first-class citizens임을 보여준다. 또한 'calc' 함수에 뺄셈을 위한 익명 함수를 직접 전달한다.
// 유형 문을 사용한 함수 프로토타입 정의: calculator 유형은 두 개의 정수를 취하고 정수를 반환하는 함수 프로토타입으로 정의된다. 이 유형은 'calc' 함수의 매개변수 유형으로 사용되며, 더욱 깔끔하고 재사용 가능한 코드를 위해 함수 프로토타입을 정의하고 사용하는 방법을 보여준다.

// english
// Anonymous Function: We define an anonymous function that calculates the sum of a variadic number of integers and assign it to the variable sum. This function is then called with a series of integers as its arguments.
// First-class Function: The add function is defined as a variable and passed as an argument to the calc function, showcasing that functions are first-class citizens in Go. We also directly pass an anonymous function for subtraction to the calc function.
// Function Prototype Definition Using Type Statement: The calculator type is defined as a function prototype that takes two integers and returns an integer. This type is used as the parameter type for the calc function, illustrating how you can define and use function prototypes for cleaner and more reusable code.
