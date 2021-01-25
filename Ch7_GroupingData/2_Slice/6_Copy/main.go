package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := make([]int, 10)
	//目標 s2 來源s1
	copy(s2, s1)
	fmt.Printf("s2:%v\n", s2)
	//會複製一份新的s2 修改s2不引響s1
	s2[0] = 100
	fmt.Printf("s1:%v\n", s1)
	fmt.Printf("s2:%v\n", s2)
}
