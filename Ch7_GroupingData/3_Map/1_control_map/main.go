package main

import "fmt"

func main() {

	stm := map[string]int{
		"Ken":   80,
		"Lindy": 70,
	}
	fmt.Println(stm)
	fmt.Println(stm["Lindy"])
	v, ok := stm["Vivin"]
	fmt.Println(v)
	fmt.Println(ok)
	if v, ok := stm["Lindy"]; ok {
		fmt.Println(v)
	}

	//map使用range取值
	for key, value := range stm {
		fmt.Println(key, value)
	}
	//使用make建立map
	map2 := make(map[string]int)
	map2["Lucy"] = 95
	map2["Ken"] = 95
	score, ok := map2["Lucy"]
	fmt.Println(score, ok)

}
