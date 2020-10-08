package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//一次性讀取檔案所有內容到byte[]
	//如果檔案過大會很佔記憶體
	b, err := ioutil.ReadFile("c:/MyDir/msg.txt")
	if err != nil {
		fmt.Println(err)
	}
	msg := string(b)
	fmt.Println(msg)
}
