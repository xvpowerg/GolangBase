package main

import "fmt"

func sum(value ...int) int {
	sum := 0
	for _, v := range value {
		sum += v
	}
	return sum
}

func main() {
	v1 := sum(1, 2, 3, 4, 5)
	fmt.Println(v1)
	values := []int{5, 6, 7, 8, 9}
	v2 := sum(values...)
	fmt.Println(v2)

}