package main

import (
	"fmt"
	"strconv"
)

func main() {

	//基本類型轉字串
	//參數 1 int64 位元的整數 2 10進位
	/*var v1 int32 = 20
	str1 := strconv.FormatInt(int64(v1), 10)
	fmt.Println(str1)
	var x2 float32 = 123.456789
	//參數 1 轉換為字串的數值 2 轉換後顯示的模式 3 取道小數點第幾位 4 轉換的變數是幾位元
	s32 := strconv.FormatFloat(float64(x2), 'f', 2, 32)
	fmt.Println(s32)*/

	//字串轉浮點數 如果有錯誤err不會是nil
	str2 := "13.56"
	f1, err := strconv.ParseFloat(str2, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(f1)

	//字串轉整數
	//參數1  要轉的文字 2 表示10進位 3 表示轉為64位元
	str3 := "250"
	v2, _ := strconv.ParseInt(str3, 10, 64)
	fmt.Println(v2)

	//字串轉[]byte
	str4 := "Howard"
	b := []byte(str4)
	fmt.Println(b[0])
	//byte轉字串 也可以放一個Slice
	fmt.Println(string(b[1]))
	//整數轉字串
	str5 := strconv.Itoa(1234)
	fmt.Println(str5)
}
