package main

import (
	"fmt"
	"log"
	"os"
	"sync"
)

func testLongTime() {
	for i := 1; i <= 100000; i++ {
		log.Println(i)
	}
	wg.Done()
}

var wg sync.WaitGroup

func main() {
	//測試會跑很久
	//testLongTime()
	//加入go 執行後很快就完畢了
	logf, err := os.OpenFile(`C:\MyDir\golog.log`,
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0666)
	if err != nil {
		fmt.Println("err:", err)
	}
	go testLongTime()
	wg.Add(1)
	log.SetOutput(logf)
	log.Println("=======完成!========")
	wg.Wait()
}
