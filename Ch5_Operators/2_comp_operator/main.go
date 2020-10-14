package main

import "fmt"

func main() {
	v1 := 15
	v2 := 2
	fmt.Printf("%v > %v =%t\n", v1, v2, v1 > v2)
	fmt.Printf("%d < %d =%v\n", v1, v2, v1 < v2)
	fmt.Printf("%d >= %d =%v\n", v1, v2, v1 >= v2)
	fmt.Printf("%d <= %d =%v\n", v1, v2, v1 <= v2)
	fmt.Printf("%d == %d =%v\n", v1, v2, v1 == v2)
	fmt.Printf("%d != %d =%v\n", v1, v2, v1 != v2)
}
