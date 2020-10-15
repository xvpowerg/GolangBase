package main

import "fmt"

type person struct {
	id   int
	name string
}

func main() {

	p1 := person{
		id:   1,
		name: "Ken",
	}
	p2 := person{
		id:   2,
		name: "Vivin",
	}

	perArryPoint := make([]*person, 0, 5)
	perArryPoint = append(perArryPoint, &p1, &p2)

	for _, p := range perArryPoint {
		//希望Slice內的person id都改為10
		p.id = 10
	}
	//成功
	fmt.Printf("perArryPoint[0]:%v perArryPoint[1]:%v\n",
		perArryPoint[0], perArryPoint[1])
	fmt.Printf("p1:%v p2:%v\n", p1, p2)

}
