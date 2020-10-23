package main

import "fmt"

func main() {
	b1 := true
	b2 := false
	//且
	fmt.Printf("%t && %t = %t\n", b1, b2, b1 && b2)
	//或
	fmt.Printf("%t || %t = %t\n", b1, b2, b1 || b2)
	//反向
	fmt.Printf("!%t  = %t\n", b1, !b1)
}
