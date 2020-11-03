package main

import "fmt"

func main() {
	s1 := []int{6, 8, 9}
	fmt.Printf("default Slice cap:%d\n", cap(s1)) //容量
	fmt.Printf("default Slice len:%d\n", len(s1)) //長度

	s2 := make([]int, 3, 20)
	fmt.Printf("s2 len:%d\n", len(s2))  //長度
	fmt.Printf("s2  cap:%d\n", cap(s2)) //容量
	fmt.Printf("s2:%v\n", s2)           //value
	fmt.Printf("s2[5]:%d\n", s2[5])     //index超過長度

	/*arr := make([]int, 8, 10)
	fmt.Println(arr)
	fmt.Println("len:", len(arr))
	fmt.Println("cap:", cap(arr))
	fmt.Printf("原始指標:%p\n", arr)

	arr = append(arr, 16)
	fmt.Println(arr)
	fmt.Println("len:", len(arr))
	fmt.Println("cap:", cap(arr))

	arr = append(arr, 72)
	fmt.Println(arr)
	fmt.Println("len:", len(arr))
	fmt.Println("cap:", cap(arr))
	fmt.Printf("無超過容量的指標:%p\n", arr)

	arr = append(arr, 83)
	fmt.Println(arr)
	fmt.Println("len:", len(arr))
	fmt.Println("cap:", cap(arr))   //超過容量Slice會自動增加2倍
	fmt.Printf("超過容量的指標:%p\n", arr) //會產生新的slice 位置跟無超過容量的指標不一樣*/

}
