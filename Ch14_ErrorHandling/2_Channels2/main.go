package main

import (
	"fmt"
	"log"
	"os"
)

func testLongTime() {
	for i := 1; i <= 100000; i++ {
		log.Println(i)
	}
	c <- 1
}

var c = make(chan int)

func main() {
	//測試會跑很久
	//testLongTime()
	//加入go 執行後很快就完畢了
	logf, err := os.OpenFile(`C:\MyDir\golog.log`,
		os.O_CREATE|os.O_TRUNC|os.O_WRONLY,
		0666)
	if err != nil {
		fmt.Println("err:", err)
	}
	go testLongTime()
	log.SetOutput(logf)
	log.Println("=======完成!========")
	<-c //這裡會暫停

}
