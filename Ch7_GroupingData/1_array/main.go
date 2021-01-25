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
	//i 是index
	//v 是value
	for i, v := range arr2 {
		fmt.Printf("idx: %d,v: %d\n", i, v)
	}
	//可不須給長度 Go自動計算
	arra3 := [...]string{"A", "B", "C"}
	for idx2, v2 := range arra3 {
		fmt.Printf("index:%v value:%v", idx2, v2)
	}

}
