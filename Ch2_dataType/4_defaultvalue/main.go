package main

import "fmt"

func main() {
	var height int32
	fmt.Println(height)
	var weight float64
	fmt.Println(weight)
	//fmt.Printf("%.1f\n", weight)
	var isOpen bool
	fmt.Println(isOpen)
	var name string
	fmt.Println(name)
	//slices
	var values []int
	fmt.Println(values == nil)
	fmt.Println("values:", values)
	var myFun func()
	fmt.Println("myFun:", myFun)

}
