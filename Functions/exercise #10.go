package main

import "fmt"

func main() {
	f := abc()
	fmt.Println(f())
	fmt.Println(f())

	g := abc()
	fmt.Println(g())
	fmt.Println(g())
}

func abc() func() int  {
	x := 0
	return func() int{
		x++
		return x
	}
}