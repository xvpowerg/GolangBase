package main

import "fmt"

func main() {
	//複合指定運算子
	count := 0
	count += 5
	fmt.Println(count) //5
	count -= 3         //count = count - 3
	fmt.Println(count) //2
	count *= 6
	fmt.Println(count) //12
	count /= 2         //count = count / 2
	fmt.Println(count) //6

}
