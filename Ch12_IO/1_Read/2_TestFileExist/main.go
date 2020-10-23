package main

import (
	"fmt"
	"os"
)

func main() {

	f, _ := os.Open(`C:\MyDir\msg.txt`)
	_, err := f.Stat()
	if err == nil {
		fmt.Println("exist!")
	} else if os.IsNotExist(err) {
		fmt.Println("IsNotExist!")
	}
	//也可用封裝好的 os.Stat與 os.IsNotExist 搭配
	filePath := `C:\MyDir\msg.txt`
	_, err2 := os.Stat(filePath)
	fmt.Println(os.IsNotExist(err2))
	//os.IsExist 是給檔案存在 但沒權限讀取時使用

}
