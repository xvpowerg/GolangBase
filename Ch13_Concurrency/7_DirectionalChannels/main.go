package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano()) //因為要使用亂數必須給予亂數的來源
}
func testRecive(ch <-chan int) {
	fmt.Printf("v :%d", <-ch)
}

func testSend(ch chan<- int) {
	fmt.Println("資料計算中...")
	time.Sleep(time.Duration(2) * time.Second)
	fmt.Println("資料完成傳送資料...")
	ch <- rand.Int()
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
