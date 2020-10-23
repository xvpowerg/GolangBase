package main

import "fmt"

func main() {
	//Slice無須給長度
	lis := []int{5, 2, 9}
	//i為index v為slice內的數值
	for i, v := range lis {
		fmt.Println(i, v)
	}
}
