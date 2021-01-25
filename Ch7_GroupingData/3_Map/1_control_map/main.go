package main

import "fmt"

func main() {

	stm := map[string]int{
		"Ken":   80,
		"Lindy": 70,
		"Vivin": 25,
	}
	fmt.Println(stm)
	fmt.Println(stm["Lindy"])
	v, ok := stm["Vivin"]
	fmt.Println(v)
	fmt.Println(ok) //key是否存在

	if v, ok := stm["Lucy"]; ok {
		fmt.Println(v)
	}

	//map使用range取值
	/*for key, value := range stm {
		fmt.Println(key, value)
	}*/

	//使用make建立map
	/*map2 := make(map[string]int)
	map2["Lucy"] = 95
	map2["Ken"] = 78
	score, ok := map2["Lucy"]
	fmt.Println(score, ok)*/
	//使用make建立map 也可加入長度
	len := 3
	map3 := make(map[string]int, len)
	map3["A"] = 10
	map3["B"] = 20
	map3["C"] = 30
	fmt.Println(map3)

}
