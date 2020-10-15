package main

import "fmt"

type student struct {
	id    int
	name  string
	score float32
}

func main() {
	st1 := student{
		1,
		"Ken",
		72.0,
	}
	fmt.Println(st1)
	//推薦以下方式
	st2 := student{
		name:  "Vivin",
		id:    2,
		score: 72.0,
	}
	fmt.Println(st2)

}
