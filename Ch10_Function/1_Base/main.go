package main

import (
	"fmt"
	"strings"
)

func test1() {
	fmt.Println("Hello!! Method!")
}
func test2(s string) {
	fmt.Println("msg:", s)
}
func test3(s string) string {
	return strings.ToUpper(s)
}
func test4(s string) (string, int) {
	return strings.ToLower(s), len(s)
}
func main() {
	test1()
	test2("Gigi")
	name := test3("iris")
	fmt.Println(name)
	name, strlen := test4("LUCY")
	fmt.Println(name, strlen)

}
