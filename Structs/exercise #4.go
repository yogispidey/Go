package main

import "fmt"

func main() {
	s := struct {
		first map[int]string
		second []string
	}{
		first: map[int]string{
			1: "yogi",
			2: "spider",
		},
		second: []string{"yogi","man"},
	}

	fmt.Println(s)
}
