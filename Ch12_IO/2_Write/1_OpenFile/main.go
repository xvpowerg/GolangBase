package main

import (
	"fmt"
	"os"
)

func main() {
	/*
	   O_RDONLY  // open the file read-only.只能讀
	   O_WRONLY // open the file write-only.只能寫
	   O_RDWR    // open the file read-write.能讀寫
	   O_APPEND // append data to the file when writing.附加文字
	   O_CREATE // create a new file if none exists. 如果檔案不存在建立
	   O_EXCL    // used with O_CREATE, file must not exist. 跟O_CREATE一起使用 建立的檔案不可存在
	   O_SYNC    // open for synchronous I/O.同步IO
	   O_TRUNC   // truncate regular writable file when opened.會清空開啟的檔案
	*/
	//0666 表示 可讀可寫
	//Linux
	f, e := os.OpenFile(`C:\MyDir\data.txt`, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer f.Close()
	if e != nil {
		fmt.Println(e)
		return
	}
	f.WriteString("\n")
	rn, _ := f.WriteString("AAA")
	rn2, _ := f.WriteString("BBBCC") //rn,rn2寫入文字數量

	fmt.Println(rn, rn2)

	//使用以下方式也可寫入字串 效果一樣
	//io.WriteString(f, "Howard")
	//io.WriteString(f, "Ken")

}
