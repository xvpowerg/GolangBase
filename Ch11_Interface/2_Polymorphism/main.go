package main

import "fmt"

type Square struct {
	height int
	width  int
}
type Circle struct {
	radius float64
}

type CalculateArea interface {
	area()
}

func (t Square) area() {
	fmt.Printf("Square: %d\n", t.height*t.width)
}
func (c Circle) area() {
	fmt.Printf("Product: %.2f\n", c.radius*3.1415)
}

func testArea(calArea CalculateArea) {
	calArea.area()
}

func main() {
	s := Square{
		height: 10,
		width:  20,
	}
	c := Circle{
		radius: 10,
	}

	testArea(s)
	testArea(c)

}
