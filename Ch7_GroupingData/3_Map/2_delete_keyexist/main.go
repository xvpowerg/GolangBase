package main

import "fmt"

func main() {

	map2 := make(map[string]int)
	map2["Lucy"] = 95
	map2["Ken"] = 72
	fmt.Println(map2)
	if _, ok := map2["Lucy"]; ok {
		delete(map2, "Lucy")
	}
	fmt.Println(map2)
}
