package main

import "fmt"

func test1() {
	c := make(chan int, 3)
	//三輸入 不用等待
	c <- 7
	c <- 3
	c <- 6
	//c <- 7 //會出錯因為輸出3次了
	fmt.Println(<-c) //輸出7
	fmt.Println(<-c) //輸出3
	fmt.Println(<-c) //輸出6
	//可以用for接收chan嗎?可以
}

//使用for接收channel
func test2() {
	c := make(chan int, 3)
	go func() {
		for i := 1; i <= 3; i++ {
			c <- i
		}
		//close(c)
	}()
	for {
		if v, ok := <-c; ok {
			fmt.Println("v:", v)
		} else {
			break
		}

	}

}

func main() {
	test1()
	//test2()
}
