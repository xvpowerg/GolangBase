package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var testNumber = 3
var c = make(chan int)

func sendMsg(i int) {
	fmt.Printf("傳送i:%d中資料...\n", i)
	c <- i
	fmt.Printf("傳送i:%d傳送完成...\n", i)
}

func receive() {
	for v := range c {
		//fmt.Printf("接收v:%d資料等待2秒...\n", v)
		time.Sleep(time.Duration(2) * time.Second)
		fmt.Printf("接收v:%d...\n", v)
		wg.Done()
	}
}

func main() {
	for i := 1; i <= testNumber; i++ {
		wg.Add(1)
		go sendMsg(i)
	}
	go receive()
	wg.Wait()
	close(c)
}
