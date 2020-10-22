package main

import (
	"errors"
	"fmt"
)

/*
type error interface {
    Error() string
}
*/

func division(a int32, b int32) (int32, error) {
	res := a / b
	var err error
	if b == 0 {
		err = errors.New("分母不可為0")
	}
	return res, err
}
func main() {
	v1, err := division(10, 2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v1)

}
