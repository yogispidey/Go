package main

import "fmt"

func main() {
	x := foo()
	y, z := bar()
	fmt.Println(x, y, z)
}

func foo() int {
	return 0
}

func bar() (int, string) {
	return 1, "yogi"
}
