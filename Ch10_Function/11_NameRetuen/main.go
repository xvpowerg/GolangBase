package main

import "fmt"

//無錄影
//定義name return
//可指定return變數名稱
//好處不需要特別定義回傳數值
func test1(v1 int, v2 int) (x1 int, x2 int) {
	x1 = v1 + 10
	x2 = v2 - 20
	return
}
func main() {
	y1, y2 := test1(10, 60)
	fmt.Printf("a=%v b=%v\n", y1, y2)
}
