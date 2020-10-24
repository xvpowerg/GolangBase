package main

import "fmt"

func receive1(evenChan, oddChan <-chan int, quit <-chan bool) ([]int, []int) {
	evenList := make([]int, 0)
	oddList := make([]int, 0)
	for {
		select {
		case even := <-evenChan:
			evenList = append(evenList, even)
		case odd := <-oddChan:
			oddList = append(oddList, odd)
		case <-quit:
			return evenList, oddList
		}

	}

}

func send1(even, odd chan<- int, quit chan<- bool) {
	for i := 1; i <= 100; i++ {
		switch {
		case i%2 == 0:
			even <- i
		default:
			odd <- i
		}
	}
	close(quit)
}

func main() {

	evenChan := make(chan int)
	oddChan := make(chan int)
	quit := make(chan bool)
	go send1(evenChan, oddChan, quit)
	evenList, oddList := receive1(evenChan, oddChan, quit)
	fmt.Println(evenList, oddList)

}
