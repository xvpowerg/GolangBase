package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//注意屬性必須是大寫開頭
type product struct {
	Id    int
	Name  string
	Price float32
}

func main() {
	jsonData, _ := ioutil.ReadFile("product.json")
	var p product
	//注意必須是point
	json.Unmarshal(jsonData, &p)
	fmt.Println(p)

}
