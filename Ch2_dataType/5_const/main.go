package main

import "fmt"

func main() {

	const PI float64 = 3.1415927
	const (
		a int     = 42
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
	fmt.Println(Wed)

	//宣告Animal
	type Animal int
	type Lnage int
	const (
		Tiger Animal = iota
		Cat
		Dog
	)
	const (
		C Lnage = iota
		Go
		Java
	)
	fmt.Println(Tiger)
	fmt.Println(Cat)
	var zoo Animal = Tiger
	fmt.Println(zoo)
	//java不是Lnage 類型所以錯誤
	//var zoo2 Animal = Java

}
