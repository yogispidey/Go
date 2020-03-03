package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64
	//var v int
	var gr int = 6
	var wg sync.WaitGroup

	wg.Add(gr)

	for i := 0; i < gr; i++ {
		go func() {
			atomic.AddInt64(&counter,1)
			//runtime.Gosched()
			//v := counter
			//runtime.Gosched()
			//v++
			//counter = v
			fmt.Println("value:", atomic.LoadInt64(&counter))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("final value:", counter)
}