package main

import (
	"fmt"
)

type rect struct {
	width, heigh int
}

func (r *rect) area() int {
	return r.width * r.heigh
}

func (r *rect) perim() int {
	return 2 * (r.heigh + r.width)
}

func main() {
	r := rect{10, 12}
	fmt.Println(r.area())
	fmt.Println(r.perim())
}
