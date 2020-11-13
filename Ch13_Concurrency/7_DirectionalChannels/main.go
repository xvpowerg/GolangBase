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

	c := make(chan int)
	cr := make(<-chan int) //recive
	cs := make(chan<- int) //send

	fmt.Printf("c %T\n", c)
	fmt.Printf("cr %T\n", cr)
	fmt.Printf("cs %T\n", cs)

	go testSend(c)
	testRecive(c)
}
