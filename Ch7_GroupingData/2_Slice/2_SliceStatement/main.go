package main

import "fmt"

func main() {
	//修改產生的Slice會影響來源
	values := []int{81, 92, 63, 74, 45, 66, 57, 18, 19}
	fmt.Println("default:", values)
	newValues := values[1:]
	newValues[2] = 77
	fmt.Println("newValues:", newValues)
	fmt.Println("default:", values)

	fmt.Println(values[1:3])
	end := 4
	fmt.Println(values[1:end])

	//array to Slice
	/*arr := [3]int{1, 2, 3}
	slice1 := arr[0:]
	fmt.Printf("arr type%T\n", arr)
	fmt.Printf("slice1 type%T\n", slice1)*/

	//不可負數
	//fmt.Println(values[-1:-3])
	//不可大到小
	//fmt.Println(values[3:1])

	//加入邊界測試
	//fmt.Println(values[0:2:5])
	//fmt.Println(values[0:6:5]) //錯誤因為結尾長度大於邊界
}
