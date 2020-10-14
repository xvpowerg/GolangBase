package main

import "fmt"

func main() {
	//Slice無須給長度
	lis := []int{5, 2, 9}

	for i, v := range lis {
		fmt.Println(i, v)
	}
}
