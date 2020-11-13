package main

import (
	"fmt"
)

var count int = 10
var c = make(chan int)

func sendMsg() {
	for i := 1; i <= count; i++ {
		c <- i
	}
	//close(c)
}
func receive() {
	for v := range c {
		fmt.Println("v:", v)
	}
}
func main() {
	go sendMsg()
	receive()

}
