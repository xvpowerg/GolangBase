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

func (t Test1) copy() Test1 {
	return t
}

func main() {
	t1 := Test1{
		name: "Vivin",
	}
	t1.SetName("Ken")
	fmt.Printf("t1:%v\n", t1)
	t1.SetNamePoint("Lucy")
	fmt.Printf("t1:%v\n", t1)

	t2 := t1.copy()
	t2.name = "Iris"
	fmt.Printf("t1:%v\n", t1)
	fmt.Printf("t2:%v\n", t2)

}
