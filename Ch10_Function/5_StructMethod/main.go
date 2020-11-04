package main

import "fmt"

type person struct {
	id   int
	name string
}

//Typte模式
func (p person) info() string {
	return fmt.Sprintf("id:%d name:%s\n", p.id, p.name)
}

//Point模式
func (p *person) infoP() string {
	return fmt.Sprintf("id:%d name:%s\n", p.id, p.name)
}
func main() {
	pe1 := person{id: 10,
		name: "Ken"}
	fmt.Println(pe1.info())
	fmt.Println(pe1.infoP())

}
