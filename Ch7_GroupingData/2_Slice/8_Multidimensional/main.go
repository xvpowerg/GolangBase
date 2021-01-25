package main

import "fmt"

func main() {
	//多維陣列
	array2x3 := [2][3]int{{1, 2, 3},
		{4, 5, 6}}
	fmt.Println("array2x3[0][2]:", array2x3[0][2])
	//輪巡多維陣列
	for _, array := range array2x3 {
		for _, v2 := range array {
			fmt.Printf("%v ", v2)
		}
		fmt.Println()
	}
	fmt.Println("==================")
	//多維Slice
	//建立一組長度為2的Slice 並且此Slice可以放置一組int的slice
	slice2x2 := make([][]int, 2)
	//於Slice內附加Slice
	slice2x2[0] = []int{1, 2}
	slice2x2[1] = []int{5, 6}
	for _, slice2 := range slice2x2 {
		for _, v := range slice2 {
			fmt.Printf("%v ", v)
		}
		fmt.Println()
	}
	//多維Slice建立法2 建立一組2*3的陣列
	slice2x3 := make([][]int, 2)
	for i := range slice2x3 {
		slice2x3[i] = make([]int, 3)
	}

}
