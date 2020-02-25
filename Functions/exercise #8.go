package main

import "fmt"

func abc() func() int  {
	return func() int {
		return 3
	}
}

func main() {
	x := abc()
	fmt.Println(x())
}
