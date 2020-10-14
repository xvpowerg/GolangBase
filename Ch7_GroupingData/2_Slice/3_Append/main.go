package main

import "fmt"

func main() {
	values := []int{1, 2, 3, 4}
	fmt.Println(values)
	values2 := append(values, 5, 6)
	fmt.Println("values2:", values2)
	//合併Slice
	values3 := []int{50, 90, 80}
	values4 := append(values, values3...)
	fmt.Println("values4:", values4)
	//array to Slice
	arr := [3]int{1, 2, 3}
	slice1 := arr[0:]
	fmt.Printf("arr type%T\n", arr)
	fmt.Printf("slice1 type%T\n", slice1)

}
