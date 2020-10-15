package main

import "fmt"

func main() {
	a1 := 10
	b2 := 20

	b2 = a1
	b2 = 95
	fmt.Printf("a1:%d b2:%d\n", a1, b2)
	fmt.Println("======使用Point==========")
	p1 := &a1
	*p1 = 95
	var p2 *int = &a1
	//%p 可輸出16進位的Point
	//p1內存放Point
	//*p1取p1內存放Point的數值
	fmt.Printf("p1:%p *p1:%d a1:%d\n", p1, *p1, a1)
	fmt.Printf("p2:%p *p2:%d a1:%d\n", p2, *p2, a1)

}
