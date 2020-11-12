package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file, err := os.Open(`C:\MyDir\msg.txt`)
	defer file.Close()
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	//接收一組Reader 用途方便讀取資料不用自己跑回圈
	//適合小檔案 因為Buffer不能調整
	//Buffer 預設為512byte
	b, e := ioutil.ReadAll(file)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Printf("%s", b)
}
