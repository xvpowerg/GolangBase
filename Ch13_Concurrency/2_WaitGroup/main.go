package main

import (
	"fmt"
	"log"
	"os"
	"sync"
)

func testLongTime() {
	//跑100萬次迴圈
	for i := 1; i <= 1000000; i++ {
		log.Println(i)
	}
	wg.Done() //計數器減1
}

var wg sync.WaitGroup

func main() {

	//加入go 執行後很快就完畢了
	logf, _ := os.OpenFile(`C:\MyDir\golog.log`,
		os.O_CREATE|os.O_TRUNC|os.O_WRONLY,
		0666)
	log.SetOutput(logf)
	go testLongTime()
	wg.Add(1) //計數器加1
	fmt.Println("=======完成!========")
	wg.Wait() //wg 計數器到0為止都會等待
}
