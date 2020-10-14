package main

import "fmt"

func main() {
	values := []int{1, 2, 3, 4}
	fmt.Println(values)
	values = append(values[0:1], values[2:]...)
	fmt.Println(values)
}
