package main

import "fmt"

func receive1(c1, c2 <-chan int, quit <-chan bool) {

	for {
		select {
		case v := <-c1:
			fmt.Printf("c1 value:%d\n", v)
		case v := <-c2:
			fmt.Printf("c2 value:%d\n", v)
		case <-quit:
			return
		}

	}

}

func send1(c1, c2 chan<- int, quit chan<- bool) {
	c2 <- 20
	c1 <- 10
	close(quit)
}

func main() {

	c1 := make(chan int)
	c2 := make(chan int)
	quit := make(chan bool)
	go send1(c1, c2, quit)
	receive1(c1, c2, quit)

}
