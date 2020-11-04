package main

import "fmt"

func main() {
	type person struct {
		id   int
		name string
	}

	p1 := person{
		id:   1,
		name: "Ken",
	}
	p2 := person{
		id:   2,
		name: "Vivin",
	}

	perSlicePoint := make([]*person, 0, 5)
	perSlicePoint = append(perSlicePoint, &p1, &p2)

	for _, p := range perSlicePoint {
		//希望Slice內的person id都改為10
		p.id = 10
	}
	//成功
	fmt.Printf("perSlicePoint[0]:%v perSlicePoint[1]:%v\n",
		perSlicePoint[0], perSlicePoint[1])
	fmt.Printf("p1:%v p2:%v\n", p1, p2)

}
