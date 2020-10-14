package main

import "fmt"

func main() {
	age := 18
	if age >= 18 {
		fmt.Println("成年")
	} else {
		fmt.Println("未成年")
	}
	//可以使用初始化指令
	if y := 2 + 5; y >= 7 {
		fmt.Println(">=7")
	}
}
