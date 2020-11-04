package main

import "fmt"

func main() {
	type person struct {
		id   int
		name string
	}
	//無使用指標
	p1 := person{
		id:   1,
		name: "Ken",
	}
	p2 := p1 //將p1複製了一份到p2
	p2.id = 30
	fmt.Println("p1:", p1)
	fmt.Println("p2:", p2)

	//使用Struct指標
	/*p3 := person{
		id:   1,
		name: "Ken",
	}
	p4 := &p3 //將p3的Point放到p4
	//var p4 *person = &p3 //將p1的Point放到p3
	//注意不須使用*p4
	p4.id = 75
	fmt.Printf("p3:%v\n", p3)
	fmt.Printf("p4:%v\n", p4)*/

	//實務上寫法如下=======================
	/*p5 := &person{
		id:   1,
		name: "Ken",
	}
	p6 := p5
	p6.name = "Iris"
	fmt.Printf("p5:%v", p5)*/

	//	不推薦以下寫法
	/*	p5 := new(person)
			p5.id = 25
			p5.name = "Vivin"
		p6 := p5
		p6.name = "Iris"
		fmt.Printf("p5:%v", p5)*

	*/

}
