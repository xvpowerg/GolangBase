package main

import "fmt"

type Test1 struct {
	name string
}

func (t Test1) SetName(n string) {
	t.name = n
}
func (t *Test1) SetNamePoint(n string) {
	t.name = n
}

func main() {
	t1 := Test1{
		name: "Vivin",
	}
	t1.SetName("Ken")
	fmt.Println(t1)
	t1.SetNamePoint("Lucy")
	fmt.Println(t1)
}
