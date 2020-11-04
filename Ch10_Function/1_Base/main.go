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

//傳入二組變數，使用return回傳一組變數
func test3(s1 string, s2 string) string {
	return strings.ToUpper(s1) + strings.ToUpper(s2)
}

//傳入一組變數，回傳組兩組變數 ()記得加
func test4(s string) (string, int) {
	return strings.ToLower(s), len(s)
}
func main() {
	test1()
	//test2("Gigi")

	//name := test3("iris", "vivin")
	//fmt.Println(name)

	//接兩組變數的方式
	//name, strlen := test4("LUCY")
	//fmt.Println(name, strlen)

}
