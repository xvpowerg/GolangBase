//package 是資料夾名稱
package obj

import (
	"fmt"
)

type Dog struct {
	Age  int
	Name string
}

func (d Dog) info() {
	fmt.Printf("Name:%s Age:%d\n", d.Name, d.Age)
}
