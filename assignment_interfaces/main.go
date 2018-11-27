package main

import (
	"fmt"
)

type triangle struct {
	base   float64
	height float64
}

type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func main() {
	triangle := triangle{base: 10, height: 5}
	square := square{sideLength: 7}

	fmt.Print("Square Area: ")
	printArea(square)
	fmt.Print("Triangle Area: ")
	printArea(triangle)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}
