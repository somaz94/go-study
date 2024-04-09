package main

import "fmt"

func maps() {
	// Declare and initialize a map using make
	idMap := make(map[int]string)
	// Add or update map entries
	idMap[901] = "Apple"
	idMap[134] = "Grape"
	idMap[777] = "Tomato"

	// Declare and initialize a map using literals
	tickers := map[string]string{
		"GOOG": "Google Inc",
		"MSFT": "Microsoft",
		"FB":   "FaceBook",
		"AMZN": "Amazon",
	}

	// Reading and printing a value from idMap
	fmt.Println("134:", idMap[134])

	// Checking for a key's existence in tickers and printing
	val, exists := tickers["MSFT"]
	if exists {
		fmt.Println("MSFT:", val)
	} else {
		fmt.Println("No MSFT ticker")
	}

	// Deleting an entry from idMap
	delete(idMap, 777)

	// Enumerating over tickers with a for range loop
	fmt.Println("Tickers:")
	for key, val := range tickers {
		fmt.Println(key, val)
	}

	// Enumerating over idMap with a for range loop
	fmt.Println("idMap:")
	for key, val := range idMap {
		fmt.Println(key, val)
	}
}

// korean
// 맵 선언 및 초기화: make 함수와 리터럴 구문을 각각 사용하여 idMap과 tickers라는 두 개의 맵을 생성한다.
// 항목 추가 및 업데이트: idMap에 새 키-값 쌍을 추가하고 기존 키에 새 값을 할당하여 기존 쌍을 업데이트한다.
// 항목 읽기: 키를 지정하여 idMap에서 값을 읽고 인쇄합니다. 키가 존재하지 않으면 Go는 맵의 값 유형에 대해 0 값을 반환한다.
// 키 존재 확인: 티커에 특정 키가 존재하는지 확인하고 존재하는 경우 그 값을 인쇄하는 방법을 보여준다.
// 항목 삭제: delete 기능을 사용하여 idMap에서 키-값 쌍을 제거한다.
// 맵 열거: 'for range' 루프를 사용하여 'idMap'과 'tickers'를 모두 반복하여 모든 키-값 쌍을 인쇄한다. 맵에 대한 반복 순서는 지정되지 않으며 실행마다 다를 수 있다.

// english
// Map Declaration and Initialization: We create two maps, idMap and tickers, using both the make function and literal syntax, respectively.
// Adding and Updating Entries: We add new key-value pairs to idMap and update existing ones by simply assigning a new value to an existing key.
// Reading Entries: We read and print a value from idMap by specifying its key. If a key doesn't exist, Go returns the zero value for the map's value type.
// Checking Key Existence: We demonstrate how to check whether a specific key exists in tickers and print its value if it does.
// Deleting Entries: We remove a key-value pair from idMap using the delete function.
// Enumerating Maps: We use a for range loop to iterate over both idMap and tickers, printing all key-value pairs. The iteration order over maps is not specified and can vary between executions.
