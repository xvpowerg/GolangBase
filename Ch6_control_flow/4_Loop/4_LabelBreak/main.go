package main

import "fmt"

func main() {
Loop:
	for i := 1; i <= 3; i++ {
		fmt.Printf("Start %d\n", i)
		for k := 1; k <= 5; k++ {
			fmt.Printf("%d*%d=%2d ", i, k, i*k)
			if i == 2 {
				//break Loop
				continue Loop
			}
		}
		fmt.Println()
		fmt.Printf("End%d \n", i)
	}

}
