package main

import "fmt"

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
func main() {
	f1 := incrementor()
	f2 := incrementor()

	f1()
	f1()
	f1()
	f2()
	fmt.Println(f1())
	fmt.Println(f2())
}
