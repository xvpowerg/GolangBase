package main

import "fmt"

func main() {
	//Slice無須給長度
	lis := []int{5, 2, 9}
	//i為index v為slice內的數值
	for i, v := range lis {
		fmt.Println(i, v)
	}

	//字串轉[]byte
	str4 := "Howard"
	b := []byte(str4)
	for i, v := range b {
		fmt.Println(i, v)
	}
	fmt.Println(string(b[1]))
}
