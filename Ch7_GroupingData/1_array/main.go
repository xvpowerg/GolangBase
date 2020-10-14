package main

import "fmt"

func main() {
	//array有給長度
	arr := [6]int{4, 5, 6, 7, 8, 9}
	fmt.Println(arr[2])

	var arr2 [3]int
	arr2[0] = 10
	arr2[1] = 51
	arr2[2] = 32

	fmt.Println("len:", len(arr2))

	for i, v := range arr2 {
		fmt.Printf("idx: %d,v: %d\n", i, v)
	}

}
