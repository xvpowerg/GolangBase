package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//使用檔案名稱讀檔
	//Buffer會設定為檔案得長度+512
	b, err := ioutil.ReadFile(`c:\MyDir\msg.txt`)
	if err != nil {
		fmt.Println(err)
	}
	msg := string(b)
	fmt.Println(msg)
}
