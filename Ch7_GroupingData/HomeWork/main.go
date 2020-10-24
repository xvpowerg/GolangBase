package main

import "fmt"

//宣告一組陣列 內容如下 5 8 10 1 13 7 將陣列內容加總
func test1() {
	arr := [6]int{5, 8, 10, 1, 13, 7}
	sum := 0
	for _, v := range arr {
		sum += v
	}
	fmt.Println("sum:", sum)
}

//宣告兩組陣列 內容如下 陣列1:99,76,88,32 陣列2: 1,2,3,4,5,6,7,8,9 將兩陣列合併
func test2() {
	arr1 := [4]int{99, 76, 88, 32}
	arr2 := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	newArray := append(arr1[0:], arr2[0:]...)
	fmt.Println("newArray:", newArray)
}

//宣告一組Slice 內容如下 9,1,7,4,5,2 幫我把 7 與 5 delete
func test3() {
	list := []int{9, 1, 7, 4, 5, 2}
	list = append(list[0:2], list[3], list[5])
	fmt.Println("list:", list)
}

//宣告一組Slice 內容如下 Ken,Vivin,Lindy,Joy,Iris
//幫我分解成2組Slice[ Ken,Vivin,Lindy] 與[Joy,Iris]
func test4() {
	list := []string{"Ken", "Vivin", "Lindy", "Joy", "Iris"}
	s1 := list[0:3]
	s2 := list[3:]
	fmt.Println("s1:", s1, "s2:", s2)
}
func main() {
	test4()
}
