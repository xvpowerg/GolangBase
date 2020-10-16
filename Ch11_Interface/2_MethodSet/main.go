package main

import "fmt"

type item struct {
	id   int
	name string
}
type PrintItem interface {
	printId()
	printName()
}

func (i item) printId() {
	fmt.Println(i.id)
}
func (i *item) printName() {
	fmt.Println(i.name)
}

func printItem(pri PrintItem) {
	pri.printId()
	pri.printName()
}

func main() {
	// T  可看到 T的方法
	//*T  可看到 T的方法與*T的方法

	it := item{
		id:   10,
		name: "Ken",
	}
	it.printId()
	it.printName()
	//printItem(it)//錯誤因為it不是point所以缺少printName
	printItem(&it)

}
