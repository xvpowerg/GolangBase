package main

import (
	"fmt"
	"strings"
)

//無傳入無回傳
func test1() {
	fmt.Println("Hello!! Method!")
}

//有傳入無回傳
func test2(s string) {
	fmt.Println("msg:", s)
}

//有傳入回傳一組數值
func test3(s1 string, s2 string) string {
	return strings.ToUpper(s1) + strings.ToUpper(s2)
}

//有傳入回傳組兩組數值
func test4(s string) (string, int) {
	return strings.ToLower(s), len(s)
}

func init() {
	fmt.Println("init() 只會運行一次 可做初始化用")
}
func main() {
	test1()
	test2("Gigi")
	name := test3("iris", "vivin")
	fmt.Println(name)
	name, strlen := test4("LUCY")
	fmt.Println(name, strlen)

}
