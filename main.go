package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	r.height++
	return r.width * r.height
}

func (r rect) perim() int {
	r.height--
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}
	fmt.Println("r:", r)

	fmt.Println("area: ", r.area())
	fmt.Println("r:", r)
	fmt.Println("perim:", r.perim())
	fmt.Println("r:", r)

	rp := &r
	fmt.Println("rp:", rp)
	fmt.Println("area: ", rp.area())
	fmt.Println("rp:", rp)
	fmt.Println("perim:", rp.perim())
	fmt.Println("rp:", rp)
}
