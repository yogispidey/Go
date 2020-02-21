package main

import "fmt"

func main() {
	xi := []string {
		"James", "Bond", "Shaken, not stirred",
	}
	fmt.Println(xi)
	xj := [] string {
		"Miss", "Moneypenny", "Helloooooo, James.",
	}
	fmt.Println(xj)
	xx := [][]string {
		xi,xj,
	}
	fmt.Println(xx)

	for i:=0; i<len(xx); i++{
		for j:=0; j<len(xx[i]); j++{
			fmt.Printf("Position is %d and value is %s\n", j, xx[i][j])
		}
	}

	for i,val := range xx {
		fmt.Printf("Record:%d\n",i)
		for j,v := range val {
			fmt.Printf("\tPostion is: %d\tValue is: %s\n",j,v)
		}
	}
}
