package main

import (
	"fmt"
	"sort"
)

/*
type Interface interface {
    // Len is the number of elements in the collection.
    Len() int
    // Less reports whether the element with
    // index i should sort before the element with index j.
    Less(i, j int) bool
    // Swap swaps the elements with indexes i and j.
    Swap(i, j int)
}
*/
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
	peoples := []Person{
		{"Ken", 45},
		{"Lucy", 31},
		{"Iris", 19},
		{"Vivin", 26},
	}

	fmt.Println(peoples)
	//把people 轉型為 ByAge
	//https://golang.org/pkg/sort/#Sort
	sort.Sort(ByAge(peoples))
	fmt.Println(peoples)

}
