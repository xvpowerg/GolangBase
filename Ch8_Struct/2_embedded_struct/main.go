package main

import "fmt"

type book struct {
	isbn  string
	name  string
	price int
}
type student struct {
	id    int
	name  string
	score float32
	book
}

func main() {
	//嵌入
	st2 := student{
		name:  "Vivin",
		id:    2,
		score: 72.0,
	}

	st2.book = book{
		isbn:  "AB0001",
		name:  "Android",
		price: 50,
	}

	fmt.Println(st2)

	st3 := student{
		name:  "Vivin",
		id:    2,
		score: 72.0,
		book: book{
			isbn:  "AC0001",
			name:  "IOS",
			price: 65,
		},
	}
	fmt.Println(st3)

}
