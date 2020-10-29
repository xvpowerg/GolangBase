package main

import "fmt"

func main() {
	b1 := true
	b2 := true
	//且 兩邊為真才為真 身高165以上 且 托福500以上 才可報名
	fmt.Printf("%v && %v = %v\n", b1, b2, b1 && b2)
	//或 單邊為真就是真 軍警票或學生票都可8折
	fmt.Printf("%t ||  %t = %t\n", b1, b2, b1 || b2)
	//反向 唱反調
	fmt.Printf("!%t  = %t\n", b1, !b1)
}
