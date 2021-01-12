package main

import "fmt"

/*
無錄影
閉包
*/
func AddUpper() func(int) int {
	var n int = 10
	return func(x int) int {
		n = n + x
		return n
	}
}
func main() {
	f1 := AddUpper()
	//可以把回傳的函數當成物件 n當作回傳函數的屬性
	//第一次n回傳11因為10+1
	fmt.Println(f1(1)) //輸出11
	//第二次n回傳13因為11+2
	fmt.Println(f1(2)) //輸出13

}
