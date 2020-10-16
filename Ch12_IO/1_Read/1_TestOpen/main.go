package main

import (
	"fmt"
	"os"
)

func readFile() {
	//只可讀取因為權限為O_RDONLY
	file, err := os.Open(`C:\MyDir\msg.txt`)
	defer file.Close()
	if err != nil {
		fmt.Println("file error:", err)
		return
	}
	buffArray := make([]byte, 4096)

	for {
		index, err := file.Read(buffArray)
		if err != nil {
			fmt.Println(err)
			break
		}
		//fmt.Printf("v:%v", byteArray[:index])
		fmt.Printf("v:%s", buffArray[:index])
	}
}
func main() {
	readFile()
}
