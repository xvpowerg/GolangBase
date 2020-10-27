package main

import "fmt"

func main() {
	//type 關鍵字表示宣告某種類型
	//DataInfo 自定義的類型名稱
	//int32 表示DataInfo 可存放的內容
	type DataInfo int32
	var data1 DataInfo = 10
	//%d 表示整數 %T 表示類型
	fmt.Printf("%d %T", data1, data1)
	type Student string

}
