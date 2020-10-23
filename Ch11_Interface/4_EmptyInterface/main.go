package main

import "fmt"

func main() {
	//可以是任意類型
	var v1 interface{}
	var v2 interface{}
	v1 = 10
	v2 = 20.56
	fmt.Println(v1, v2)
	//強制轉型(Type Assertion) 只能用於interface
	m, ok := v1.(int)
	if !ok {
		fmt.Println("轉型錯誤!")
	}
	fmt.Println(m, ok)
	//可使用switch輸出類型
	switch t := v2.(type) {
	case int:
		fmt.Println("int", t)
	case float64:
		fmt.Println("float", t)
	default:
		fmt.Println("Error")

	}

}
