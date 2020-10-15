package main

import "fmt"

type person struct {
	id   int
	name string
}

func main() {
	p1 := person{
		id:   1,
		name: "Ken",
	}

	p2 := p1 //將p1複製了一份到p2
	p2.id = 30

	fmt.Println("p1:", p1)
	fmt.Println("p2:", p2)

	//p3:= &p1
	var p3 *person = &p1 //將p1的Point放到p3
	p3.id = 75
	fmt.Println("p1:", p1)
	fmt.Println("p3:", *p3)

}
