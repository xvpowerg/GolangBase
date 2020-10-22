package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var testNumber = 5000
var c = make(chan int)

func sendMsg(i int) {
	c <- i
}

func receive() {
	for {
		//如果耗時在這家buffer也沒用
		//time.Sleep(time.Duration(1) * time.Millisecond)
		if _, ok := <-c; !ok {
			break

		}
		wg.Done()
	}
}

func main() {
	t1 := time.Now() // get current time
	for i := 1; i <= testNumber; i++ {
		wg.Add(1)
		go sendMsg(i)
	}
	go receive()
	wg.Wait()
	close(c)
	elapsed := time.Since(t1)
	fmt.Printf("elapsed:%v\n", elapsed)
}
