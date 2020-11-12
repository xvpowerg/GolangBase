package main

import (
	"fmt"
	"os"
)

func readFile() {
	//https://golang.org/pkg/io/#Reader
	//只可讀取因為權限為O_RDONLY
	file, err := os.Open(`C:\MyDir\msg.txt`)
	defer file.Close()
	if err != nil {
		fmt.Println("file error:", err)
		return
	}
	buffArray := make([]byte, 4096)

	for {

		end, err := file.Read(buffArray)
		//io.EOF 表示資料讀完了
		if err != nil {
			fmt.Println(err)
			break
		}
		//fmt.Printf("v:%v", byteArray[:index])
		fmt.Printf("v:%s", buffArray[:end])
	}
}
func main() {
	readFile()
}
