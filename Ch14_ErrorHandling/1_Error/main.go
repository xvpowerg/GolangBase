package main

import (
	"errors"
	"fmt"
)

func division(a int32, b int32) (int32, error) {
	var err error
	if b == 0 {
		err = errors.New("分母不可為0") //常見建立error的方式呼叫 errors.New
		return 0, err
	}
	res := a / b

	return res, nil
}
func main() {
	v1, err := division(10, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v1)

}
