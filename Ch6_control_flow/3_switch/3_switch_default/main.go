package main

import "fmt"

func main() {

	age := 18
	switch {
	case age < 18:
		fmt.Println("少年")
	case age <= 60:
		fmt.Println("青年")
	case age <= 80:
		fmt.Println("壯年")
	default:
		fmt.Println("老年")
	}

}
