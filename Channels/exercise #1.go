package main

import (
	"fmt"
)

func main() {
	c := make(chan int,1)
	//works with buffered channel
	c <- 42
	//works
	/*go func() {
		c <- 42
	}()*/

	fmt.Println(<-c)
}
