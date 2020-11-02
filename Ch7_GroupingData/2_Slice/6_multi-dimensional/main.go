package main

import "fmt"

func main() {
	//長度3 容量5
	value1 := make([]int, 3, 5)
	//長度6 容量10
	value2 := make([]int, 6, 10)
	lis2d := [][]int{value1, value2}
	lis2d[0][1] = 10
	fmt.Println(lis2d)
	fmt.Println(lis2d[0][1])

	//二維Slice輪巡
	for _, v1 := range lis2d {
		for _, v2 := range v1 {
			fmt.Print(v2, " ")
		}
		fmt.Println()
	}

}
