package main

import "fmt"

//可傳入方法到另一個方法
func calculat(x int, y int, calcu func(int, int) int) int {
	return calcu(x, y)
}
func main() {
	f1 := func(x int, y int) int {
		return x + y
	}
	f2 := func(x int, y int) int {
		return x - y
	}
	fmt.Println(calculat(5, 6, f1))
	fmt.Println(calculat(8, 9, f2))

}
