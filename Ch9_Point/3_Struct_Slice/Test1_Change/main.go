package main

import "fmt"

func main() {
	type person struct {
		id   int
		name string
	}
	perSlice := make([]person, 0, 5)
	p1 := person{
		id:   1,
		name: "Ken",
	}
	p2 := person{
		id:   2,
		name: "Vivin",
	}
	perSlice = append(perSlice, p1, p2)
	//per是複製一份person
	per := perSlice[0]
	//修改id
	per.id = 12
	//是否影響Slice呢?
	fmt.Printf("perSlice%v\n", perSlice)
	fmt.Printf("p1:%v\n", p1)
	fmt.Printf("p1:%p per:%p\n", &p1, &per)

}
