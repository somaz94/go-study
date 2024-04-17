package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// say prints a string several times.
func printRepeatedly(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s, "***", i)
		time.Sleep(time.Millisecond * 500) // simulate a task
	}
}

func goroutines() {
	// Set the number of CPUs to use.
	runtime.GOMAXPROCS(4)

	// Create WaitGroup to wait for multiple goroutines.
	var wait sync.WaitGroup
	wait.Add(3) // We will wait for three goroutines.

	// execute function synchronously
	printRepeatedly("Sync")

	// Execute the function asynchronously using goroutines
	go func() {
		defer wait.Done()
		printRepeatedly("Async1")
	}()
	go func() {
		defer wait.Done()
		printRepeatedly("Async2")
	}()
	go func() {
		defer wait.Done()
		printRepeatedly("Async3")
	}()

	// Use anonymous functions with goroutines
	go func(msg string) {
		defer wait.Done()
		fmt.Println(msg)
	}("Running anonymous function")

	// Wait for all goroutines to complete
	wait.Wait()
	fmt.Println("All goroutines finished executing")
}

// korean
// Go의 동시성과 병렬성 개념은 Go의 성능 특성과 여러 작업을 처리하는 방법의 핵심이다. 이러한 개념에 대한 설명과 함께 Go에서 고루틴, 익명 함수 및 CPU 활용의 실제 적용을 보여주는 예를 제공한다.
// goroutine: 함수 호출 전에 go 키워드를 사용하면 printRepeatedly 함수와 익명 함수가 비동기적으로 실행됩니다. 이를 통해 서로 동시에 그리고 주 프로그램과 함께 실행할 수 있습니다.
// sync.WaitGroup: 고루틴 그룹의 실행이 완료될 때까지 기다리는 데 사용됩니다. 'Add'는 대기할 고루틴 수를 설정하고, 'Done'은 활성 고루틴의 카운터를 감소시키며, 'Wait'은 모든 고루틴이 'Done'을 호출할 때까지 차단합니다.
// runtime.GOMAXPROCS: 동시에 실행될 수 있는 최대 CPU 수를 설정합니다. 이를 통해 Go는 진정한 병렬 실행을 위해 여러 코어를 활용할 수 있으며 이는 CPU 바인딩 작업에 유용합니다.
// 동시성 대 병렬성:
// 동시성은 많은 일을 한 번에 처리하는 것입니다(프로그램을 독립적으로 실행되는 단위로 구성). 이 예시에서는 여러 고루틴이 작업을 동시에 처리합니다.
// 병렬성은 여러 작업을 동시에 수행하는 것입니다. 예에서 'runtime.GOMAXPROCS'를 통해 활성화되며 잠재적으로 여러 CPU 코어에서 이러한 작업을 동시에 실행할 수 있습니다.

// english
// This Go program demonstrates goroutines, anonymous functions, sync.WaitGroup, and setting the number of logical CPUs to handle concurrency and parallelism.
// Goroutines: Using the go keyword before function calls, the printRepeatedly function and an anonymous function are executed asynchronously. This allows them to run concurrently with each other and with the main program.
// sync.WaitGroup: This is used to wait for a group of goroutines to finish executing. Add sets the number of goroutines to wait for, Done decrements the counter of active goroutines, and Wait blocks until all goroutines have called Done.
// runtime.GOMAXPROCS: This sets the maximum number of CPUs that can be executing simultaneously. This allows Go to utilize multiple cores for true parallel execution, which is useful for CPU-bound tasks.
// Concurrency vs. Parallelism:
// Concurrency is about dealing with many things at once (structuring your program as independently executing units). In this example, multiple goroutines handle tasks concurrently.
// Parallelism is about doing many things at the same time, enabled by runtime.GOMAXPROCS in the example, which potentially runs these tasks on multiple CPU cores simultaneously.
