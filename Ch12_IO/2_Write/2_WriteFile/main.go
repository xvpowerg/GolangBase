package main

import (
	"io/ioutil"
)

func main() {
	//WriteFile 內包裝了OpenFile 讓我們方便寫檔案
	var size uint = 1024 * 1024 * 300 //300mb
	data := make([]byte, size, size)  //
	ioutil.WriteFile(`C:\MyDir\TestData.txt`, data, 0666)
	//把字串轉換為byte Slice
	strData := []byte("ABCDEFGHIJKLMNOPQR")
	ioutil.WriteFile(`C:\MyDir\TestDataStr.txt`, strData, 0666)

}
