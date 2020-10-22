package main

import "fmt"

func main() {
	for x := 1; x <= 10; x++ {
		if x == 6 {
			break
		}
		fmt.Print(x, " ")
	}
	fmt.Println()
	for x := 1; x <= 10; x++ {
		if x%2 == 0 {
			continue
		}
		fmt.Print(x, " ")
	}

}
