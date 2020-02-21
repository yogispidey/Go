package main

import "fmt"

type person struct {
	first string
	last string
	favourite []string
}

func main() {
	p1 := person {
		first: "yogi",
		last: "yogi",
		favourite: []string{"chocolate","almond"},
	}

	p2 := person{
		first:     "spider",
		last:      "man",
		favourite: []string{"butterscotch", "blueberry"},
	}

	fmt.Println(p1.first)
	fmt.Println(p1.last)

	for i,v := range p1.favourite {
		fmt.Println(i,v)
	}

	fmt.Println(p2.first)
	fmt.Println(p2.last)

	for i,v := range p2.favourite {
		fmt.Println(i,v)
	}

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _,v := range m{
		fmt.Println(v.last)
		for i,val := range v.favourite{
			fmt.Println(i,val)
		}
	}
}
