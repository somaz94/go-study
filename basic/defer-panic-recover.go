package main

import (
	"fmt"
	"os"
)

func defer_panic_recover() {
	// Attempt to open a file and handle errors and panics
	openFile("1.txt") // You might change this filename to an invalid one to see panic handling
	fmt.Println("File handling done. Continue with other tasks...")

	// Incorrect file name entered to demonstrate panic and recover
	openFile("Invalid.txt")

	// Due to recovery, this statement will be executed even after panic
	fmt.Println("Done")
}

// openFile tries to open a file and read from it
func openFile(fn string) {
	// defer function. Executed when panic is called
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	// Ensure file is closed after all operations are done
	defer f.Close()

	// read file
	bytes := make([]byte, 1024)
	n, err := f.Read(bytes)
	if err != nil {
		panic(err) // triggers the deferred panic recovery
	}
	fmt.Println("Bytes read:", n)
}

// korean
// 기능: 이 코드는 파일 이름으로 지정된 파일을 열고 읽으려고 시도한다. 패닉 복구 메커니즘을 통해 잠재적인 오류를 처리한다.
// defer: openFile 함수가 정상적으로 종료되었는지 패닉으로 인해 종료되었는지 여부에 관계없이 openFile 함수 실행이 완료되면 파일이 닫히도록 한다.
// panic: 파일을 열거나 읽는 동안 오류가 발생하면 함수에 패닉이 발생하고, 패닉에서 복구하도록 설계된 지연된 함수에 의해 이를 포착한다.
// recover: 패닉을 포착하여 프로그램이 계속 실행되도록 한다. main 함수의 마지막 인쇄 문에서 알 수 있듯이 패닉을 일으킨 오류를 인쇄하고 정상적인 실행을 다시 시작한다.

// english
// Functionality: This code tries to open and read from a file specified by the filename. It handles potential errors through panic-recover mechanism.
// defer: Ensures the file is closed when the function openFile finishes execution, regardless of whether it finishes normally or due to a panic.
// panic: If an error occurs during file opening or reading, the function panics, which is then caught by the deferred function designed to recover from panics.
// recover: Captures the panic, allowing the program to continue executing. It prints the error that caused the panic and resumes normal execution, as evidenced by the final print statement in the main function.
