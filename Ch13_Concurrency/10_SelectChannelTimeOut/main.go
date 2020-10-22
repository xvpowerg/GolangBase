package main

import (
	"fmt"
	"time"
)

func receive1(c1, c2 <-chan int) {
	for {
		select {
		case v := <-c1:
			fmt.Printf("c1 value:%d\n", v)
		case v := <-c2:
			fmt.Printf("c2 value:%d\n", v)
		case <-time.After(1 * time.Second): //一秒後沒收到訊息自動結束
			fmt.Println("After Stop ")
			return
		}
	}

}

func send1(c1, c2 chan<- int) {
	c2 <- 20
	c1 <- 10
}

func main() {

	c1 := make(chan int)
	c2 := make(chan int)

	go send1(c1, c2)
	receive1(c1, c2)

}
