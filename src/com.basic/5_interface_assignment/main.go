package main

import (
	"fmt"
)

type shape interface {
	getArea() float64
	printArea()
}

type traingle struct {
	height, base float64
}

type square struct {
	sideLength float64
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t traingle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) printArea() {
	fmt.Println("Area is :", s.getArea())
}
func (t traingle) printArea() {
	fmt.Println("Area is :", t.getArea())
}

func main() {
	sSquare := square{sideLength: 3}
	sTraingle := traingle{height: 2, base: 4}

	sSquare.printArea()
	sTraingle.printArea()
}
