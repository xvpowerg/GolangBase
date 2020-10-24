package main

import "fmt"

//使用迴圈顯示2~50當中的所有的質數，每5筆斷行
//數字1不算質數
func test1() {
	for i, count := 2, 1; i <= 50; i++ {
		k := 0
		for k = i - 1; k > 1; k-- {
			//只要碰到能整除了就離開迴圈
			if i%k == 0 {
				break
			}
		}
		//除到最後只有1能整除了，他就是質數
		if k == 1 {
			fmt.Printf("%2d ", i)

			if count%5 == 0 {
				fmt.Println()
			}
			count++
		}
	}
}

//使用迴圈計算費氏數列12是多少?
func test2() {
	//因為費氏數列第0項與第1項是數學定義，
	//費氏數列第0項為0
	//費氏數列第1項為1
	//所以計算費氏數列12是<12，因為第一項不用算，已經定義好了
	f1, f2 := 0, 1
	for i := 1; i < 12; i++ {
		f1, f2 = f2, f1+f2
	}
	fmt.Println("f2:", f2)
}
func main() {
	//test1()
	test2()
}
