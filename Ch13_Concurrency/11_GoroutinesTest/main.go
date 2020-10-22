package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func sum(count int) {
	sum := 0
	for i := 1; i <= count; i++ {
		sum += i
	}
	fmt.Printf("Sum:%d\n", sum)
	wg.Done()
}

func main() {
	now := time.Now()
	wg.Add(2)
	go sum(900000000)
	go sum(900000000)
	wg.Wait()
	elapsed := time.Since(now)
	fmt.Printf("elapsed:%v\n", elapsed)

}
