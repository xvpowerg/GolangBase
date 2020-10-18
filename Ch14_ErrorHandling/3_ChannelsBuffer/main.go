package main

import "fmt"

func main() {
	c := make(chan int, 3)
	//三輸入 不用等待
	c <- 7
	c <- 3
	c <- 6
	//c <- 7 //會出錯因為輸出3次了
	fmt.Println(<-c) //輸出7
	fmt.Println(<-c) //輸出3
	fmt.Println(<-c) //輸出6
	//可以用for接收chan嗎?

}
