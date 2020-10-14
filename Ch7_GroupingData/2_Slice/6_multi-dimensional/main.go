package main

import "fmt"

func main() {
	//長度5 容量10
	value1 := make([]int, 5, 10)
	//長度10 容量20
	value2 := make([]int, 10, 20)
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
