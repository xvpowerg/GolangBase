package main

import "fmt"

func main() {
	price := 10
	fmt.Printf("價格:%d\n", price)
	height := 156.12678
	//取道小數點第二位
	fmt.Printf("身高:%.2f\n", height)

	name := "Ken"
	weight := 56.234
	//輸出字串
	str := fmt.Sprintf("姓名:%s 體重:%.2f \n", name, weight)
	fmt.Printf(str)
	//輸出指標類型
	str = fmt.Sprintf("Point %p \n", &name)
	fmt.Printf(str)
	//輸出name的類型
	str = fmt.Sprintf("Type %T \n", name)
	fmt.Printf(str)
	//自動判斷name的類型給預設的格式
	str = fmt.Sprintf("vale %v \n", name)
	fmt.Printf(str)
}
