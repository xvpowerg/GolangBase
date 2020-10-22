package main

import "fmt"

func main() {

	for i := 1; i <= 9; i++ {
		for k := 1; k <= 9; k++ {
			fmt.Printf("%d*%d=%2d ", i, k, i*k)
		}
		fmt.Println()
	}
}
