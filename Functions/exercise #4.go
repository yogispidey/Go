package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println(p.first, p.last, p.age)
}

func main() {
	p1 := person{
		first: "yogi",
		last:  "spider",
		age:   24,
	}
	p1.speak()
}
