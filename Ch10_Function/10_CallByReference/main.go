package main

import "fmt"

type Test1 struct {
	name string
}

func swap(a *int, b *int) {
	*b, *a = *a, *b
}
func changeName(v *Test1) {
	v.name = "Lindy"
}

func changeSliceValue(args []int) {
	fmt.Printf("args:%p\n", args)
	args[1], args[0] = args[0], args[1]
}
func main() {
	a := 1
	b := 65
	swap(&a, &b)
	fmt.Printf("a:%d,b:%d\n", a, b)

	t1 := Test1{
		"Ken",
	}
	changeName(&t1)
	fmt.Println(t1)

	//Slice 預設為CallByReference
	values := []int{15, 27}
	fmt.Printf("values:%p\n", values)
	fmt.Println(values)
	changeSliceValue(values)
	fmt.Println(values)
}
