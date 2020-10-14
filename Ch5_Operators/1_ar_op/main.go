package main

import "fmt"

func main() {
	//算數運算子
	v1 := 15
	v2 := 2
	v4 := 15.0
	v3 := 2.0
	ans1 := v1 + v2
	fmt.Printf("%d + %d=%d\n", v1, v2, ans1)
	ans2 := v1 - v2
	fmt.Printf("%d - %d=%d\n", v1, v2, ans2)
	ans3 := v1 * v2
	fmt.Printf("%d * %d=%d\n", v1, v2, ans3)
	ans4 := v1 / v2
	fmt.Printf("%d / %d=%d\n", v1, v2, ans4)
	ans5 := v1 % v2
	fmt.Printf("%d %% %d=%d\n", v1, v2, ans5)
	ans6 := v4 / v3
	fmt.Printf("%.2f / %.2f=%.2f\n", v4, v3, ans6)
}
