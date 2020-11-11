package main

import "fmt"

func swap(x *int, y *int) {
	*x, *y = *y, *x
}

func main() {
	a := 1
	b := 65
	swap(&a, &b)
	fmt.Printf("a:%d,b:%d\n", a, b)
}
