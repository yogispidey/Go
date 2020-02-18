package main

import "fmt"

type yogi int
var x yogi
var y int

var z=10

func main() {

	fmt.Println(x)
	fmt.Printf("%T\n",x)
	x=42
	fmt.Println(x)
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T",y)

}
