package main

import "fmt"

func testDefer() {
	defer fmt.Println("testDefer")
	fmt.Println("Test2")
	fmt.Println("Test3")
}

func testDeferOrder() {
	defer fmt.Println("Test1_1")
	fmt.Println("Test2")
	defer fmt.Println("Test1_2")
	fmt.Println("Test3")
	defer fmt.Println("Test1_3")
}

func testVar(i int) {
	//i變數會被儲存
	defer fmt.Println("TestVar Setp1:", i)
	i++
	defer fmt.Println("TestVar Setp2:", i)
}

func main() {
	testDefer()
	//testDeferOrder()
	//testVar(1)
}
