package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func receive2(c1, c2 <-chan int, quit <-chan bool) {

	for {
		select {
		case v := <-c1:
			fmt.Printf("c1 value:%d\n", v)
		case v := <-c2:
			fmt.Printf("c2 value:%d\n", v)
		case <-quit:
			return
		}
		wg.Done()
	}

}

func send2(c1, c2 chan<- int) {
	go func() {
		for i := 1; i <= 5; i++ {
			c1 <- i

		}
	}()

	go func() {
		for i := 6; i <= 10; i++ {
			c2 <- i
		}
	}()

}

func main() {
	wg.Add(10)
	c1 := make(chan int)
	c2 := make(chan int)
	q := make(chan bool)
	send2(c1, c2)
	go receive2(c1, c2, q)
	wg.Wait()
	q <- true

}
