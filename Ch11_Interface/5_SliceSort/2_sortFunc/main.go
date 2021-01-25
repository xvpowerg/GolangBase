package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

func main() {
	peoples := []Person{
		{"Ken", 45},
		{"Lucy", 31},
		{"Iris", 19},
		{"Vivin", 26},
	}
	//小到大排序
	sortFunc := func(i, k int) bool {
		return peoples[i].Age < peoples[k].Age
	}
	//SliceStable Sort
	//https://golang.org/pkg/sort/#example_SliceStable
	sort.SliceStable(peoples, sortFunc)
	fmt.Println(peoples)

	//Sort Map
	sortMap := make(map[int]string)
	keySlice := make([]int, 0)
	sortMap[10] = "Ken"
	sortMap[1] = "Vivin"
	sortMap[2] = "Lindy"
	sortMap[7] = "Joy"
	sortMap[8] = "Lucy"
	for k := range sortMap {
		keySlice = append(keySlice, k)
	}
	sort.Slice(keySlice, func(i, j int) bool {
		return keySlice[i] < keySlice[j]
	})
	for _, k := range keySlice {
		fmt.Printf("key:%v value:%v\n", k, sortMap[k])
	}

}
