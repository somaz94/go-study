package main

import "fmt"

func printConstantsAndVariables() {
	// Declaring a constant
	const constantValue = "immutable"
	fmt.Println("Constant:", constantValue)

	// Declaring a variable
	var variableValue = "mutable"
	fmt.Println("Variable before change:", variableValue)

	// Changing the value of the variable
	variableValue = "changed value"
	fmt.Println("Variable after change:", variableValue)
}
