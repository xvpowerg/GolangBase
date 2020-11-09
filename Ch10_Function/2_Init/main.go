package main

import (
	"fmt"
)

func init() {
	fmt.Println("init() 可做初始化用 1")
}

func main() {
	fmt.Println("main.....")
}
func init() {
	fmt.Println("init() 可做初始化用 2")
}
