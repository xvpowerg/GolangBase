package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	//WriteFile 內包裝了OpenFile 讓我們方便寫檔案
	//1024 byte = 1Kb
	//1024 Kb = 1Mb
	var size uint = 1024 * 1024 * 300 //300mb
	data := make([]byte, size, size)  //
	error1 := ioutil.WriteFile(`C:\MyDir\TestData.txt`, data, 0666)
	if error1 != nil {
		fmt.Println(error1)
		return
	}
	//把字串轉換為byte Slice
	strData := []byte("ABCDEFGHIJKLMNOPQR")
	ioutil.WriteFile(`C:\MyDir\TestDataStr.txt`, strData, 0666)

}
