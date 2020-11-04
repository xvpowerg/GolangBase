package main

import "fmt"

func main() {

	func() {
		fmt.Println("Anonymous_func!!")

	}()

	func(x int) {
		fmt.Println("Anonymous func!!", x)
	}(52)

	f := func(x int, y int) {
		fmt.Println(x, y)
	}
	f(2, 6)

}
