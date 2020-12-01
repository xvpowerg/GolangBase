package main

import "fmt"

type Person interface {
	GetName() string
}

type Student struct {
	name string
	age  int
}

func (p Student) GetName() string {
	return p.name
}
func (p Student) GetAge() int {
	return p.age
}

type Teacher struct {
	name string
	year int
}

func (p Teacher) GetName() string {
	return p.name
}
func (p Teacher) GetYear() int {
	return p.year
}

type Dog struct {
	name string
	age  int
}

func main() {

	//也可以轉其他Sturct
	var p1 Person
	//p1 = Student{name: "Ken", age: 10}
	p1 = Teacher{name: "Vivin", year: 5}
	/*s, ok2 := p1.(Student)
	if ok2 {
		fmt.Println(s.GetName(), s.GetAge())
	}*/

	//fmt.Println(p1.(type))

	//可使用switch輸出類型並且轉型
	switch t := p1.(type) {
	case Student:
		fmt.Println("Student", t.GetName(), t.GetAge())
	case Teacher:
		fmt.Println("Teacher", t.GetName(), t.GetYear())
	default:
		fmt.Println("Error")
	}

}
