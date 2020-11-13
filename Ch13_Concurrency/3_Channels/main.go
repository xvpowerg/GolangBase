package main

import (
	"fmt"
	"time"
)

func test1() {
	//預設不加數量就為0
	c := make(chan int, 0)
	c <- 1 //因為Buffer為0 所以立刻block 送
	<-c    //收
}

func test2() {
	//預設不加數量就為0
	c := make(chan int)
	go func() {
		//程式停2秒 模擬計算
		time.Sleep(time.Duration(2) * time.Second)
		c <- 100
	}()
	fmt.Println("Start...")
	v := <-c //等待接收
	fmt.Println("End...:", v)
}

func main() {
	//test1()
	test2()
}
