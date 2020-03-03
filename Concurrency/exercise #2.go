package main

import "fmt"

type person struct {
	lang string
}

func (p *person) speak()  {
	fmt.Println("He speaks",p.lang)
}

type human interface {
	speak()
}

func saySomething(h human)  {
	h.speak()
}

func main() {
	p1 := person{lang:"tamil"}
	saySomething(&p1)
	//saySomething(p1)
	//p1.speak()
}
