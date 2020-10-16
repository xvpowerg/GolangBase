package main

import "fmt"

type person struct {
	id   int
	name string
}

func (p person) info() string {

	return fmt.Sprintf("id:%d name:%s\n", p.id, p.name)
}

func main() {
	pe1 := person{id: 10,
		name: "Ken"}
	fmt.Println(pe1.info())

}
