package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	atr2 := "Hello台灣"
	//預設情況 atr2[i] 取得的數值是uint8
	//但我們使用的是16bit的數值 所以無法存所有的數字
	for i := 0; i < len(atr2); i++ {
		fmt.Printf("輸出%T\n", atr2[i])
	}
	//可將字串轉換輸出為int32位元,此時就可包含所有位元
	value3 := []int32(atr2)
	fmt.Println(value3)
	for i := 0; i < len(value3); i++ {
		fmt.Printf("輸出:%c\n", value3[i])
	}
	//字串轉整數 如果有錯誤err不等於nil
	intValue, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println(intValue)
	}

	//字串轉byte[]
	var bytes = []byte("hello go")
	fmt.Printf("byte:%v\n", bytes)
	//byte[]轉字串
	str := string(bytes)
	fmt.Printf("str:%v\n", str)
	//可輸出對應10進位的其他進位的字串 注意回傳是字串
	str2 := strconv.FormatInt(123, 16)
	fmt.Printf("str2:%v\n", str2)
	//字串是否在目前字串中 如有回傳true 否回傳false
	isFind := strings.Contains("Howard", "ow")
	fmt.Printf("isFind:%v\n", isFind)
	//回傳有幾個不重複的B
	num := strings.Count("ABBBACC", "B")
	fmt.Printf("num:%v\n", num)

	//字串比較可使用== 但是分大小寫 使用EqualFold比較不分大小寫
	b := strings.EqualFold("Iris", "iRis")
	fmt.Printf("b=%v\n", b)
	//s 表示被換的原始字串
	//old 表示被取代的字串
	//new 要取代的字
	//n表示取代個數-1表示全部取代
	rs := strings.Replace("test value", "value", "測試", 1)
	fmt.Printf("rst:%v\n", rs)
	//使用,分割
	strArr := strings.Split("A,B,C,D", ",")
	fmt.Printf("strArr:%v\n", strArr)
	//去左右2邊除空格
	trimSpaceStr := strings.TrimSpace(" Hi Go ")
	fmt.Printf("trimStr:%v\n", trimSpaceStr)
	//可指定左右兩邊要去除的內容
	trimStr := strings.Trim("! hello! ", " !")
	fmt.Printf("trimStr:%v\n", trimStr)
	//判斷字串是否為某字串開頭
	isHttps := strings.HasPrefix("https://www.google.com", "https")
	fmt.Printf("isHttps:%v\n", isHttps)

}
