package main

import "fmt"

func main() {
	str := "abcdefg"
	//想顯示"def"
	fmt.Printf("str:v1:%v\n", str[3:6])
	//因為字串不可修改所以無法使用Slice修改
	//str[1] ='B'
	//必須先轉成byte的slice才可修改
	strArr := []byte(str)
	strArr[1] = 'B'
	newStr := string(strArr)
	fmt.Println(newStr)
	//但是注意中文有問題 因為我們的byte是8bit中文是16bit
	//解法必須使用[]rune
	//type rune = int32
	str2 := "你好優"
	strArr2 := []rune(str2)
	strArr2[0] = '我'
	newStr2 := string(strArr2)
	fmt.Println(newStr2)

}
