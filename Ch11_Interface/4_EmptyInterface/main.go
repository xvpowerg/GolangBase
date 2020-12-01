package main

import "fmt"

//介面內沒有任何函數就是Empty interface
type Test interface {
}

func main() {
	//Empty interface 可以是任意類型
	var v1 Test
	var v2 interface{}
	v1 = 10
	v2 = 20.56
	fmt.Println(v1, v2)
	//使用Interface的轉型方式
	m, pass := v1.(int)
	if pass {
		fmt.Println("轉型成功!")
	} else {
		fmt.Println("轉型失敗!")
	}
	fmt.Println(m)

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
