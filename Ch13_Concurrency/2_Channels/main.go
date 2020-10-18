package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 0)
	//c <- 1 //因為Buffer為0 所以一輸入必須立刻接收到輸出
	//<-c

	go func() {
		//程式停2秒
		time.Sleep(time.Duration(2) * time.Second)
		c <- 1
	}()
	fmt.Println("Start...")
	<-c
	fmt.Println("End...")
}
