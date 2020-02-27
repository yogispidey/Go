package main

import "fmt"

type person struct {
	first string
}

func changeMe(y *person)  {
	(*y).first = "spider"
	//y.first = "spider"
}

func main() {
	x := person{first: "yogi"}
	fmt.Println(x)

	changeMe(&x)

	fmt.Println(x)
}
