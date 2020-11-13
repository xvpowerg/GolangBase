package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var testNumber = 10
var c = make(chan int, testNumber)

func sendMsg(i int) {
	c <- i
}

func receive() {
	for v := range c {
		fmt.Printf("接收v:%d資料等待2秒...\n", v)
		time.Sleep(time.Duration(2) * time.Second)
		wg.Done()
	}
}

func main() {
	//多個Goroutine的處裡方式
	for i := 1; i <= testNumber; i++ {
		wg.Add(1)
		go sendMsg(i)
	}
	go receive()
	wg.Wait()
	close(c)

}
