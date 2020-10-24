package main

import "fmt"

//製作一組函數名為reverses 可傳入Slice ，回傳兩個數值
//反轉的Slice
//此Slice的長度
//解法1
func reverses(list []int) ([]int, int) {
	reversList := make([]int, len(list))

	index := len(list) - 1
	for i, v := range list {
		reversList[index-i] = v
	}
	return reversList, len(reversList)
}

//製作一組函數名為reverses 可傳入Slice ，回傳兩個數值
//反轉的Slice
//此Slice的長度
//解法2
// func reverses(list []int) ([]int, int) {
// 	for i, k := 0, len(list)-1; i < k; i, k = i+1, k-1 {
// 		list[i], list[k] = list[k], list[i]
// 	}
// 	return list, len(list)
// }

func feb() func() int {
	var a int
	var b int = 1
	return func() int {
		ans := a + b
		a, b = b, ans
		return ans
	}
}

func main() {
	//製作一組函數名為reverses 可傳入Slice ，回傳兩個數值
	//反轉的Slice
	//此Slice的長度
	list := []int{1, 2, 3, 4, 5}
	nList, len := reverses(list)
	fmt.Println(nList, len)

	//使用閉包計算費氏數列12是多少?
	f := feb()
	ans := 0
	for i := 1; i < 12; i++ {
		ans = f()
	}
	fmt.Println("ans:", ans)
}
