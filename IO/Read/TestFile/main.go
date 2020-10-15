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
	byteArray := make([]byte, 4096)

	for {
		index, err := file.Read(byteArray)
		if err != nil {
			fmt.Println(err)
			break
		}
		//fmt.Printf("v:%v", byteArray[:index])
		fmt.Printf("v:%s", byteArray[:index])
	}
}
func main() {
	readFile()
}
