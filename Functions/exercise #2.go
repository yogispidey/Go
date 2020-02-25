package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4}
	sum := foo(xi...)
	fmt.Println(sum)

	sum2 := bar(xi)
	fmt.Println(sum2)
}

func foo(x ...int) int {
	var sum int
	for _,v := range x {
		//fmt.Println(i)
		sum += v
	}
	return sum
}

func bar(x []int) int {
	var sum int
	for _, v := range x {
		sum += v
	}
	return sum
}
