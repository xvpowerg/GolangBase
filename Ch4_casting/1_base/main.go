package main

import "fmt"

func main() {

	var v1 int32 = 20
	var v2 float32 = 25.67

	//整數轉浮點數
	v2 = float32(v1)
	fmt.Println(v2)
	//浮點數轉整數
	v1 = int32(v2)
	fmt.Println(v1)

	type MyType int
	var t1 MyType = 1
	var t2 int = int(t1)
	fmt.Println(t2)

}
