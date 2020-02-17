package main

import "fmt"

type yogi int

func main() {

	var x yogi
	fmt.Println(x)
	fmt.Printf("%T\n",x)
	x=42
	fmt.Println(x)

}
