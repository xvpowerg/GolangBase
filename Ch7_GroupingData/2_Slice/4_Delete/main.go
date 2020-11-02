package main

import "fmt"

func main() {
	values := []int{8, 7, 9, 6, 25, 31, 52}
	fmt.Println(values)
	values = append(values[0:2], values[3:]...)
	fmt.Println(values)
}
