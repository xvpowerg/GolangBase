package main

import "fmt"

func changeSliceValue(args []int) {
	fmt.Printf("args:%p\n", args)
	args[1], args[0] = args[0], args[1]
}
func main() {

	//Slice 預設為CallByReference
	values := []int{15, 27}
	fmt.Printf("values:%p\n", values)
	fmt.Println(values)
	changeSliceValue(values)
	fmt.Println(values)
}
