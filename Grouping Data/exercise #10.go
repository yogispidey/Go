package main

import "fmt"

func main() {
	favthings := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}
	favthings["yogi"] = []string{
		`cricket`,
	}
	for key, v := range favthings {
		fmt.Println(key)
		for j, val := range v {
			fmt.Printf("\tPosition is: %d\tValue is: %s\n", j, val)
		}
	}

	if _, ok := favthings["yogi"]; ok {
		delete(favthings, "yogi")
		fmt.Println("Key is deleted")
	}

	for key, v := range favthings {
		fmt.Println(key)
		for j, val := range v {
			fmt.Printf("\tPosition is: %d\tValue is: %s\n", j, val)
		}
	}
}
