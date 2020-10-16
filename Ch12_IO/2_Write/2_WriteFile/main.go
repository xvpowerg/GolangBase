package main

import (
	"io/ioutil"
)

func main() {
	var size uint = 1024 * 1024 * 300 //300mb
	data := make([]byte, size, size)  //
	ioutil.WriteFile(`C:\MyDir\TestData.txt`, data, 0666)

}
