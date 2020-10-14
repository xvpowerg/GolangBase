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
	str := fmt.Sprintf("姓名:%s 體重:%.2f \n", name, weight)
	fmt.Printf(str)
}
