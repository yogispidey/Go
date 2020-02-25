package main

import "fmt"

func abc()  {
	fmt.Println("Hi")
}

func main() {
	x := abc
	x()
}
