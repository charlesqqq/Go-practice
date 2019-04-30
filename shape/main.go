package main

import (
	"fmt"
	"math"
)

func main() {

	var s shape = square{sideLength: 5}
	fmt.Println("square:", s.getArea())
	// printArea(square{sideLength: 5})
	s = triangle{base: 10, height: 10}
	fmt.Println("triangle:", s.getArea())
	// printArea(triangle{base: 10, height: 10})
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}
type triangle struct {
	height float64
	base   float64
}

func (s square) getArea() float64 {
	return math.Pow(s.sideLength, 2)
}

func (t triangle) getArea() float64 {
	return t.base * t.height * 0.5
}
