package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var testNumber = 10
var c = make(chan int, testNumber)

func sendMsg(i int) {
	c <- i
}

func receive() {
	for v := range c {
		fmt.Printf("接收v:%d\n", v)
		//wg.Done()
	}
}

func main() {
	//多個Goroutine的處裡方式
	//wg.Add(testNumber)
	for i := 1; i <= testNumber; i++ {
		go sendMsg(i)
	}
	go receive()

	//wg.Wait()
	//close(c)

}
