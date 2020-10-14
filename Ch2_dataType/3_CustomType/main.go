package main

import "fmt"

func main() {
	type DataInfo int
	var data1 DataInfo = 10
	//%d 表示整數 %T 表示類型
	fmt.Printf("%d %T", data1, data1)

}
