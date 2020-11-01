package main

import "fmt"

func main() {
	/*for x := 1; x <= 5; x++ {
		fmt.Print(x, " ")
	}
	fmt.Println()*/

	i := 1
	//無wihle 使用for取代
	for i <= 5 {
		fmt.Print(i, " ")
		i++
	}
	fmt.Print("i:", i)
}
