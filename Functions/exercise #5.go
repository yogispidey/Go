package main

import "fmt"

type SQUARE struct {
	length float64
	width  float64
}

type CIRCLE struct {
	radius float64
}

func (s SQUARE) AREA() float64 {
	return s.length * s.width
}

func (c CIRCLE) AREA() float64 {
	return 3.14 * c.radius * c.radius
}

type SHAPE interface {
	AREA() float64
}

func INFO(s SHAPE) {
	fmt.Println("Area:", s.AREA())
}

func main() {
	c := CIRCLE{radius: 2.0}
	s := SQUARE{
		length: 2.0,
		width:  2.0,
	}
	INFO(c)
	INFO(s)
}
