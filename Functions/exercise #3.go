package main

import "fmt"

func main() {
	defer fmt.Println("this is defer, executed at the end")
	fmt.Println("This will be executed first")
}
