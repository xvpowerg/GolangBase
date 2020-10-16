package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, err := os.Open(`C:\MyDir\msg.txt`)
	defer file.Close()
	if err != nil {
		fmt.Println("file error:", err)
		return
	}
	//預設給定buff 4096byte,想改變size請使用NewReaderSize
	r := bufio.NewReaderSize(file, 1024*10) //給10kb的buffer
	//可配合Scanner使用
	i := 1
	s := bufio.NewScanner(r)
	for s.Scan() {
		fmt.Println(s.Text(), i)
		i++
	}
}
