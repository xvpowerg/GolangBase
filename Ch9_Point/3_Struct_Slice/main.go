package main

import "fmt"

type person struct {
	id   int
	name string
}

func main() {

	perArry := make([]person, 0, 5)
	p1 := person{
		id:   1,
		name: "Ken",
	}
	p2 := person{
		id:   2,
		name: "Vivin",
	}
	perArry = append(perArry, p1, p2)

	fmt.Printf("perArry[0]:%v\n", perArry[0])
	for _, p := range perArry {
		//p是複製一份person
		//我希望所有id都改為10
		p.id = 10
	}
	//結果沒變
	fmt.Printf("perArry[0]:%v perArry[1]:%v\n", perArry[0], perArry[1])
	fmt.Printf("p1:%v p2:%v\n", p1, p2)

	//per是複製一份person
	per := perArry[0]
	per.id = 12
	fmt.Println("=======================================")
	for i, _ := range perArry {
		perArry[i].id = 10
	}
	//perArry內的persion的id變為10
	fmt.Printf("perArry[0]:%v perArry[1]:%v\n", perArry[0], perArry[1])
	//結果沒變
	fmt.Printf("p1:%v p2:%v\n", p1, p2)

}
