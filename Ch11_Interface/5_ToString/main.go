package main

import "fmt"

type Item struct {
	ID    int
	Name  string
	Price int
}

/*
type Stringer interface {
    String() string
}
*/
// func (i Item) String() string {
// 	return fmt.Sprintf("ID:%d Name:%s Price:%d\n", i.ID, i.Name, i.Price)
// }
func main() {
	item := &Item{
		ID:    10,
		Name:  "Android",
		Price: 52,
	}
	fmt.Println(item)
}
