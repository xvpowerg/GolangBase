package main

import "fmt"

//只要某個Struct有print()方法，就符合PrintString interface
type PrintString interface {
	print()
}

type student struct {
	id   int
	name string
}

func (s student) print() {
	fmt.Println(s)
}

func main() {
	var pstr PrintString
	pstr = student{
		id:   10,
		name: "Ken",
	}
	pstr.print()
}
