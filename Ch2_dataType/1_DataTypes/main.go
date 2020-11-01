package main

import "fmt"

//宣告在函數外時必須使用var
var p1 = 85
var p2 int = 72
var p3 int32 = 95

func main() {

	fmt.Println(p1, p2, p3)
	//宣告在函數內可使用:=
	// x := 56
	// y := "Howard"
	// fmt.Println(x)
	// fmt.Println(y)
	// var z = 71
	// fmt.Println(z)

	// //給定類型
	//unit8
	var age byte = 255
	fmt.Println("age:", age)
	// var height float32 = 178.56
	// fmt.Println("height:", height)
	// var weight int32 = 75
	// fmt.Println("weight:", weight)

	// //複數
	// var point complex64 = 10.2 + 20.56i
	// fmt.Println(point)

	// //一次宣告多個變數
	// x1, y2 := 10, 20
	// fmt.Println(x1, y2)
	// //數字交換
	// x1, y2 = y2, x1
	// fmt.Println(x1, y2)

	// var a1, b2 int32 = 75, 63
	// fmt.Println(a1, b2)

	//bool
	var isPass bool = true
	fmt.Println(isPass)
	isOpen := false
	fmt.Println(isOpen)
	isOpen = true
	fmt.Println(isOpen)
}
