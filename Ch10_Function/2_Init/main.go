package main

import (
	"fmt"
)

func init() {
	fmt.Println("init() 只會運行一次 可做初始化用 1")
}

func init() {
	fmt.Println("init() 只會運行一次 可做初始化用 2")
}
func main() {

}
