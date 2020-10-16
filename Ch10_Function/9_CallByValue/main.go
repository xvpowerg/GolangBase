package main

import "fmt"

type Test1 struct {
	name string
}

func changeName(v Test1) {
	v.name = "Lindy"
}
func main() {
	t1 := Test1{
		"Ken",
	}
	changeName(t1)
	fmt.Println(t1)
}
