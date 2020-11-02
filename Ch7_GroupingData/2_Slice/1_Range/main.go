package main

import "fmt"

func main() {
	//Slice無須給長度
	lis := []int{5, 2, 9, 10, 15}
	fmt.Println(lis[2])
	lis[3] = 22
	//i為index v為slice內的數值
	for i, v := range lis {
		fmt.Println(i, v)
	}

	//字串轉[]byte
	/*str4 := "Howard"
	b := []byte(str4)
	for i, v := range b {
		fmt.Println(i, v)
	}
	//byte轉字串
	fmt.Println(string(b[1]))*/
}
