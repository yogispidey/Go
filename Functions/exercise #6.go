package main

import "fmt"

func main() {
	s := func(x int) int{
		return x
	}(10)
	fmt.Println(s)
}
