package main

import "fmt"

// Person struct definition
type Person struct {
	Name string
	Age  int
}

// Dict struct that contains a map
type Dict struct {
	Data map[string]string
}

// Constructor function for Dict struct
func NewDict() *Dict {
	d := Dict{}
	d.Data = make(map[string]string)
	return &d // Return pointer to Dict instance
}

func structs() {
	// Create Person instance using an empty struct and then setting fields
	var p1 Person
	p1.Name = "Bob"
	p1.Age = 20

	// Create Person instance with initial values using struct literal
	p2 := Person{Name: "Sean", Age: 50}

	// Create Person instance using new() and setting fields. All fields initialized to zero value
	p3 := new(Person)
	p3.Name = "Lee"

	// Output Person instances
	fmt.Println("Person 1:", p1)
	fmt.Println("Person 2:", p2)
	fmt.Println("Person 3:", p3)

	// Use constructor to create a Dict instance
	dic := NewDict() // Call constructor
	dic.Data["1"] = "A"

	// Output Dict instance
	fmt.Println("Dict:", dic.Data)
}

// korean
// 사람의 이름과 나이에 대한 필드를 포함하는 Person 구조체이다. 'Person'의 인스턴스는 Go에서 다양한 객체 생성 기술을 보여주기 위해 세 가지 다른 방식으로 생성된다.
// 문자열에서 문자열로의 매핑인 Data 필드를 포함하는 Dict 구조체이다. 생성자 함수 NewDict는 새로 생성된 Dict 인스턴스에 대한 포인터를 반환하기 전에 Data 필드를 빈 맵으로 초기화한다.
// 'Person' 인스턴스에 대한 포인터를 생성하기 위해 'new()' 함수를 사용하는 방법을 시연했으며, 여기에서 필드가 설정된다. Go는 포인터를 통해 구조체 필드에 액세스할 때 포인터를 자동으로 역참조하므로 구조체 포인터에서 직접 '.'을 사용할 수 있다.
// main 함수는 Person과 Dict의 인스턴스를 생성하고, 필드를 설정하고, 내용을 인쇄하는 방법을 보여준다.

// english
// A Person struct that includes fields for a person's name and age. Instances of Person are created in three different ways to demonstrate various object creation techniques in Go.
// A Dict struct that contains a Data field, which is a map from strings to strings. A constructor function, NewDict, initializes the Data field to an empty map before returning a pointer to the newly created Dict instance.
// Demonstrated use of the new() function to create a pointer to a Person instance, where fields are then set. Note that Go automatically dereferences pointers when accessing struct fields through a pointer, so you can use . directly on a struct pointer.
// The main function shows how to create instances of both Person and Dict, set their fields, and print their contents.
