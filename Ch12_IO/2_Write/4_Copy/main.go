package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	readf, err1 := os.Open(`C:\MyDir\msg.txt`)
	writef, err2 := os.OpenFile(`C:\MyDir\msg_copy.txt`,
		os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err1 != nil || err2 != nil {
		fmt.Println(err1, err2)
		return
	}
	defer readf.Close()
	defer writef.Close()
	//必須給長度
	buff := make([]byte, 8) //8bbyte buffer
	//io.Copy()//預設32*1024 = 32kb buffer
	io.CopyBuffer(writef, readf, buff)

}
