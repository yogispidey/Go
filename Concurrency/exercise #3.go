package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var counter int
	//var v int
	var gr int = 6
	var wg sync.WaitGroup

	wg.Add(gr)

	for i:= 0; i<gr; i++ {
		go func () {
			//runtime.Gosched()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			fmt.Println("value:", counter)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("final value:",counter)

	/*wg.Add(3)

	go func() {
		for i := 0;i < 10; i++ {
			v = counter
			//runtime.Gosched()
			v++
			counter=v
			fmt.Println("func1:", counter)
			//wg.Done()
		}
		wg.Done()
	}()

	go func() {
		for i := 0;i < 10; i++ {
			v = counter
			//runtime.Gosched()
			v++
			counter=v
			fmt.Println("func2:", counter)
			//wg.Done()
		}
		wg.Done()
	}()

	go func() {
		for i := 0;i < 10; i++ {
			v = counter
			//runtime.Gosched()
			v++
			counter=v
			fmt.Println("func3:", counter)
			//wg.Done()
		}
		wg.Done()
	}()

	wg.Wait()*/
}
