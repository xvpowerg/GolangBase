package main

import (
	"io/ioutil"
)

func main() {
	//將在記憶體的檔案寫出
	var size uint = 1024 * 1024 * 300 //300mb
	data := make([]byte, size, size)  //
	ioutil.WriteFile(`C:\MyDir\TestData.txt`, data, 0666)
	//把字串轉換為byte陣列
	strData := []byte("ABCDEFGHIJKLMNOPQR")
	ioutil.WriteFile(`C:\MyDir\TestDataStr.txt`, strData, 0666)

}
