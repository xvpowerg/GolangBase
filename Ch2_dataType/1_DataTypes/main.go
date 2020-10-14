package main

import "fmt"

//宣告在方法外時必須使用var
var x = 83

func main() {
	//宣告在方法內可使用:=
	x := 42 + 7
	y := "Howard"
	fmt.Println(x)
	fmt.Println(y)
	var z = 71
	fmt.Println(z)
	//給定類型
	var age byte = 100
	fmt.Println(age)
	//複數
	var point complex64 = 10 + 20i
	fmt.Println(point)
	//一次宣告多個變數
	x1, y2 := 10, 20
	fmt.Println(x1, y2)
	//數字交換
	x1, y2 = y2, x1
	fmt.Println(x1, y2)
}
