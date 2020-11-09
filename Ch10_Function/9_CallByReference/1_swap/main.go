package main

import "fmt"

func swap(a *int, b *int) {
	*b, *a = *a, *b
}

func main() {
	a := 1
	b := 65
	swap(&a, &b)
	fmt.Printf("a:%d,b:%d\n", a, b)

}
