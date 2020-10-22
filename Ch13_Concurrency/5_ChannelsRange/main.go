package main

import (
	"fmt"
)

var count int = 10
var c = make(chan int)

func sendMsg() {
	for i := 1; i <= count; i++ {
		//	fmt.Printf("傳送i:%d中資料...\n", i)
		c <- i
		//	fmt.Printf("傳送i:%d傳送完成...\n", i)
	}
	//close(c)

}

func receive() {
	for v := range c {
		//暫停1秒
		//time.Sleep(time.Duration(1) * time.Second)
		fmt.Println("v:", v)
	}
}

func main() {
	go sendMsg()
	receive()

}
