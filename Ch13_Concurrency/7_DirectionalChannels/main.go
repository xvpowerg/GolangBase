package main

import (
	"fmt"
)

func testRecive(ch <-chan int) {
	fmt.Printf("v :%d", <-ch)
}

func testSend(ch chan<- int) {
	ch <- 10
}

func main() {

	c := make(chan int)    //可相容 送與收
	cs := make(chan<- int) //send 送
	cr := make(<-chan int) //recive 收

	fmt.Printf("c %T\n", c)
	fmt.Printf("cr %T\n", cr)
	fmt.Printf("cs %T\n", cs)

	go testSend(c)
	testRecive(c)
}
