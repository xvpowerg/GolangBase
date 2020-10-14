package main

import "fmt"

func main() {
	arr := make([]int, 8, 10)
	fmt.Println(arr)
	fmt.Println(arr)
	fmt.Println(len(arr))
	fmt.Println(cap(arr))

	arr = append(arr, 16)
	fmt.Println(arr)
	fmt.Println(arr)
	fmt.Println(len(arr))
	fmt.Println(cap(arr))

	arr = append(arr, 72)
	fmt.Println(arr)
	fmt.Println(arr)
	fmt.Println(len(arr))
	fmt.Println(cap(arr))

	arr = append(arr, 83)
	fmt.Println(arr)
	fmt.Println(arr)
	fmt.Println(len(arr))
	fmt.Println(cap(arr)) //超過容量Slicp會自動增加2倍

}
