package main

import (
	"errors"
	"fmt"
)

func testSize(size int) {

	if size <= 0 {
		err := errors.New("size is Zero")
		panic(err)
	}
	fmt.Println("Pass!")
}

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err:", err)
		}
	}()
	testSize(12)

}
