//package 是資料夾名稱
package obj

import (
	"fmt"
)

//開頭小寫為私有，在相同Package能讀寫
//開頭大寫為公開，跨Package能讀寫
type Cat struct {
	age  int
	name string
}

//私有方法
func (c Cat) info() {
	fmt.Printf("Name:%s Age:%d\n", c.name, c.age)
}

//公開方法
func (c Cat) Print() {
	fmt.Printf("Name:%s Age:%d\n", c.name, c.age)
}

//公開方法
func (c *Cat) SetAge(age int) {
	c.age = age
}

//公開方法
func (c *Cat) SetName(name string) {
	c.name = name
}

//公開方法
func (c *Cat) GetAge() int {
	return c.age
}

//公開方法
func (c *Cat) GetName() string {
	return c.name
}
