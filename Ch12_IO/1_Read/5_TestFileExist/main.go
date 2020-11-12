package main

import (
	"fmt"
	"os"
)

func main() {

	// os.Stat與 os.IsNotExist 搭配查看檔案是否存在
	filePath := `C:\MyDir\msg.txt`
	_, err2 := os.Stat(filePath)
	if os.IsNotExist(err2) {
		fmt.Println("檔案不存在!")
	} else {
		fmt.Println("檔案存在!")
	}
}
