package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, e := os.OpenFile(`C:\MyDir\TestDataWrite.txt`, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer f.Close()
	if e != nil {
		fmt.Println(e)
	}
	var size uint = 1024 * 1024 * 300 //300mb
	data := make([]byte, size, size)  //產生資料
	var buffSize int = 1024 * 1024    //5kb
	w := bufio.NewWriterSize(f, buffSize)
	w.Write(data)
	w.Flush()

}
