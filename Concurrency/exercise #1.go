package main

import (
	"fmt"
	"runtime"
	"sync"
)

func foo()  {
	for i := 0; i < 10; i++{
		fmt.Println("foo:",i)
	}
	wg.Done()
}

func bar()  {
	for i := 0; i < 10; i++{
		fmt.Println("bar:",i)
	}
	wg.Done()
}

var wg sync.WaitGroup

func main() {
	fmt.Println("Inside main")
	fmt.Println("begin CPU:",runtime.NumCPU())
	fmt.Println("begin goroutine",runtime.NumGoroutine())
	wg.Add(2)
	//wg.Wait()
	go bar()
	fmt.Println("middle goroutine",runtime.NumGoroutine())
	go foo()
	fmt.Println("middle goroutine",runtime.NumGoroutine())
	wg.Wait()

	fmt.Println("end CPU:",runtime.NumCPU())
	fmt.Println("end goroutine",runtime.NumGoroutine())
}
