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
}

var wg sync.WaitGroup

func main() {

	//加入go 執行後很快就完畢了
	logf, _ := os.OpenFile(`C:\MyDir\golog.log`,
		os.O_CREATE|os.O_TRUNC|os.O_WRONLY,
		0666)
	log.SetOutput(logf)
	testLongTime()
	//go testLongTime()
	fmt.Println("=======完成!========")

}
