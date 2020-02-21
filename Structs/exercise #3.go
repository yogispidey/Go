package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct{
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	cybertrk := truck{
		vehicle:   vehicle{
			doors: 4,
			color: "red",
		},
		fourWheel: true,
	}

	verna := sedan{
		vehicle: vehicle{
			doors: 5,
			color: "blue",
		},
		luxury:  false,
	}

	fmt.Println(cybertrk,verna)
	fmt.Println(cybertrk.doors)
	fmt.Println(verna.doors)
}
