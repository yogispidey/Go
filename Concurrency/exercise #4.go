package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int
	//var v int
	var gr int = 6
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(gr)

	for i := 0; i < gr; i++ {
		go func() {
			//runtime.Gosched()
			mu.Lock()
			v := counter
			//runtime.Gosched()
			v++
			counter = v
			fmt.Println("value:", counter)
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("final value:", counter)
}