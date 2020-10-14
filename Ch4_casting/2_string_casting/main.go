package main

import (
	"fmt"
	"strconv"
)

func main() {
	var v1 int32 = 20
	//基本類型轉字串
	//參數 1 int64 位元的整數 2 10進位
	str1 := strconv.FormatInt(int64(v1), 10)
	fmt.Println(str1)

	str2 := "13.56"
	//字串轉浮點數
	f1, err := strconv.ParseFloat(str2, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(f1)
	str3 := "250"
	//字串轉整數
	//參數1  要轉的文字 2 表示10進位 3 表示轉為64位元
	v2, _ := strconv.ParseInt(str3, 10, 64)
	fmt.Println(v2)

	//字串轉[]byte
	str4 := "Howard"
	b := []byte(str4)
	fmt.Println(b[1])
	fmt.Println(string(b[1]))
}
