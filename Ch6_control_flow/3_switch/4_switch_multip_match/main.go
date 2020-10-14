package main

import "fmt"

func main() {
	name := "Ken"
	switch name {
	case "Gigi", "Lucy":
		fmt.Println("RD")
	case "Howard", "Vivin", "Ken":
		fmt.Println("經理")
	default:
		fmt.Println("Error!")
	}
}
