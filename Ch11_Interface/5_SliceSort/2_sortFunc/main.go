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
}
