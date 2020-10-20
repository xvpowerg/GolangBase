package main

import "fmt"

type Test1 struct {
	name string
}

func swap(a int, b int) {
	b, a = a, b
}
func changeName(v Test1) {
	v.name = "Lindy"
}

func main() {
	a := 1
	b := 65
	swap(a, b)
	fmt.Printf("a:%d,b:%d\n", a, b)

	t1 := Test1{
		"Ken",
	}
	changeName(t1)
	fmt.Println(t1)
}
