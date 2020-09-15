package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func newRect(width, height int) *rect {
	r := rect{width: width, height: height}
	return &r
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim: ", r.perim())

	rp := &r
	rp.height = 90
	fmt.Println("area: ", rp.area())
	fmt.Println("perim: ", rp.perim())

	newRec := newRect(5, 5)
	fmt.Println("area: ", newRec.area())
	fmt.Println("perim: ", newRec.perim())

}
