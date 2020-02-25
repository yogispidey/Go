package main

import "fmt"

func abc(f func() int) int  {
	return f()
}

func main() {
	x := func() int{
		return 20
	}
	fmt.Println(abc(x))
}
