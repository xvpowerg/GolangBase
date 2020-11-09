package main

import "fmt"

func main() {

	func() {
		fmt.Println("Anonymous_func!!")
	}()
	//匿名函數傳參數
	/*func(x int) {
		fmt.Println("Anonymous func!!", x)
	}(52)*/

	//匿名函數回傳數值
	/*v := func(x int) int {
		fmt.Println("Anonymous func!!", x)
		return x + 5
	}(52)
	fmt.Println("V:", v)*/

	//函數給變數
	/*f := func(x int, y int) {
		fmt.Println(x, y)
	}
	f(2, 6)*/

}
