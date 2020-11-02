package main

import "fmt"

func main() {
	s1 := []int{6, 8, 9}
	fmt.Printf("default Slice cap:%d\n", cap(s1)) //容量
	fmt.Printf("default Slice len:%d\n", len(s1)) //長度
	s2 := make([]int, 3, 10)
	fmt.Printf("s2 len:%d\n", len(s2)) //長度
	fmt.Printf("s2[1]:%d\n", s2[1])    //value
	fmt.Printf("s2[5]:%d\n", s2[5])    //value

	arr := make([]int, 8, 10)
	fmt.Println(arr)
	fmt.Println(len(arr))
	fmt.Println(cap(arr))
	fmt.Printf("原始指標:%p\n", arr)

	arr = append(arr, 16)
	fmt.Println(arr)
	fmt.Println(len(arr))
	fmt.Println(cap(arr))

	arr = append(arr, 72)
	fmt.Println(arr)
	fmt.Println(len(arr))
	fmt.Println(cap(arr))
	fmt.Printf("無超過容量的指標:%p\n", arr)

	arr = append(arr, 83)
	fmt.Println(arr)
	fmt.Println(len(arr))
	fmt.Println(cap(arr)) //超過容量Slice會自動增加2倍
	fmt.Printf("超過容量的指標:%p\n", arr)

}
