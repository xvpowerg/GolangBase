package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open(`C:\MyDir\test.zip`)
	wFile, err2 := os.OpenFile(`C:\MyDir\test_copy.zip`, os.O_CREATE|os.O_WRONLY, 666)
	defer file.Close()
	defer wFile.Close()

	if err != nil {
		fmt.Println("file error:", err)
		return
	}
	if err2 != nil {
		fmt.Println("file error2:", err2)
		return
	}
	byteArray := make([]byte, 4096)

	for {
		index, err := file.Read(byteArray)
		if err != nil {
			fmt.Println(err)
			break
		}
		wFile.Write(byteArray[:index])
	}

	fmt.Println(file)

}
