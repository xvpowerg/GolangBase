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

	//取得student的指標
	st3Obj := new(student)
	st3Obj.id = 1000
	st3Obj.name = "Howard"
	st3Obj.score = 85.63
	fmt.Println(st3Obj)
	//使用	st3Obj := new(student) 等同於以下語法
	// st3Obj := &student{
	// 	name:  "Howard",
	// 	id:    1000,
	// 	score: 85.63,
	// }
	// fmt.Println(st3Obj)
}
