package main

import "fmt"

func main() {

	const PI float32 = 3.1415927
	const E = 2.71828
	const PLAY = 1

	const (
		a int32   = 42
		b float64 = 42.78
		c string  = "James Bond"
	)

	//iota 會自動編號
	//宣告星期幾
	const (
		Mon = iota
		Tue
		Wed
		Thu
		Fri
		Sat
		Sun
	)
	fmt.Println(Fri)

	//宣告Animal
	type Animal int32
	type Lange int32
	const (
		Tiger Animal = iota
		Cat
		Dog
	)
	const (
		C Lange = iota
		Go
		Java
	)
	fmt.Println("Tiger:", Tiger)
	fmt.Println("Cat:", Cat)
	fmt.Println("Dog:", Dog)
	var zoo Animal = Cat
	fmt.Println(zoo)
	//java不是Lnage 類型所以錯誤
	var zoo2 Lange = Go
	fmt.Println(zoo2)

}
