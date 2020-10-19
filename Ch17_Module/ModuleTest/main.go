package main

import (
	"fmt"
	"tw/obj"
)

//myobj "tw/obj" 可使用別名
//． "tw/obj"  可以不用指定名稱

func main() {
	dog := obj.Dog{
		Age:  5,
		Name: "NiNi",
	}
	fmt.Println(dog)
}
