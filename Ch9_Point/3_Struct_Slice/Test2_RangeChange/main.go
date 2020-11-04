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
	//我希望所有id都改為10
	for _, p := range perSlice {
		//p是複製一份person
		p.id = 10
	}
	//結果沒變
	fmt.Printf("perSlice[0]:%v perSlice[1]:%v\n", perSlice[0], perSlice[1])
	fmt.Printf("p1:%v p2:%v\n", p1, p2)
}
