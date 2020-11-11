package main

import "fmt"

type Test1 struct {
	name string
}

//非point類型可用來顯示 跟修改無關的
func (t Test1) SetName(n string) {
	t.name = n
}

//point可用來顯示寫入struct 跟修改有關係的
func (t *Test1) SetNamePoint(n string) {
	t.name = n
}

func main() {
	t1 := Test1{
		name: "Vivin",
	}
	fmt.Printf("t1:%v\n", t1)
	t1.SetName("Ken")
	fmt.Printf("t1:%v\n", t1)
	t1.SetNamePoint("Lucy")
	fmt.Printf("t1:%v\n", t1)
}
