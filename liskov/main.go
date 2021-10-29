package main

import (
	"fmt"
)

func main() {
	r := Rectangle{}
	r.setHeight(5)
	r.setWidth(10)
	fmt.Println("rectangle area", r.calculateArea())

	s := Square{}
	s.setWidth(5)
	s.setHeight(4)
	fmt.Println("square area", s.calculateArea())
}