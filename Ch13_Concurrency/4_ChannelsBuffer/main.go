package main

import "fmt"

func test1() {
	c := make(chan int, 3)
	//送三組資料 不用等待
	c <- 7
	c <- 3
	c <- 6
	//c <- 9           //會出錯因為輸出3次了
	fmt.Println(<-c) //接收7
	fmt.Println(<-c) //接收3
	fmt.Println(<-c) //接收6
	//可以用for接收chan嗎?可以
}

func main() {
	test1()
}
