package main

import "fmt"

func main() {
	v1 := 15
	v2 := 2
	//大於
	fmt.Printf("%v > %v =%t\n", v1, v2, v1 > v2)
	//小於
	fmt.Printf("%d < %d =%v\n", v1, v2, v1 < v2)
	//大於等於
	fmt.Printf("%d >= %d =%v\n", v1, v2, v1 >= v2)
	//小於等於
	fmt.Printf("%d <= %d =%v\n", v1, v2, v1 <= v2)
	//等於
	fmt.Printf("%d == %d =%v\n", v1, v2, v1 == v2)
	//不等於
	fmt.Printf("%d != %d =%v\n", v1, v2, v1 != v2)
}
