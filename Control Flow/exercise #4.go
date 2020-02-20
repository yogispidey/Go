package main

import "fmt"

func main() {
	i := 1995
	for {
		if i <= 2020 {
			fmt.Println(i)
			i++
		} else {
			break
		}
	}
}
