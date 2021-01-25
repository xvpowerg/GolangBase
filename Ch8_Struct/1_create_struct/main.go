package main

import "fmt"

func main() {

	type student struct {
		id    int
		name  string
		score float32
	}

	st1 := student{
		1,
		"Ken",
		72.0,
	}
	fmt.Println(st1)
	//推薦以下方式
	st2 := student{
		id:    2,
		name:  "Vivin",
		score: 72.0,
	}
	fmt.Println(st2)
	//可透過以下方式修改
	st2.name = "Lucy"
	st2.score = 15.26
	fmt.Println(st2)

}
