package main

import (
	"fmt"
)

func main() {
	values := []int{1, 2, 3, 4, 5, 6}
	//Slice後不會改變來源 會產生新的Slice結構
	fmt.Println(values[1])
	fmt.Println(values[1:])
	fmt.Println(values[1:3])
	end := 4
	fmt.Println(values[1:end])

	//array to Slice
	arr := [3]int{1, 2, 3}
	slice1 := arr[0:]
	fmt.Printf("arr type%T\n", arr)
	fmt.Printf("slice1 type%T\n", slice1)
	//不可負數
	//fmt.Println(values[-1:-3])
	//不可大到小
	//fmt.Println(values[3:1])

}
