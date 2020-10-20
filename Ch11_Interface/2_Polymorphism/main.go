package main

import "fmt"

type Square struct {
	height int
	weight int
}
type Circle struct {
	radius int
}

type CalculateArea interface {
	area()
}

func (t Square) area() {
	fmt.Printf("Square: %d\n", t.height*t.weight)
}
func (c Circle) area() {
	fmt.Printf("Product: %.2f\n", float64(c.radius)*3.1415)
}

func testArea(calArea CalculateArea) {
	calArea.area()
}

func main() {
	s := Square{
		height: 10,
		weight: 20,
	}
	c := Circle{
		radius: 10,
	}

	testArea(s)
	testArea(c)

}
