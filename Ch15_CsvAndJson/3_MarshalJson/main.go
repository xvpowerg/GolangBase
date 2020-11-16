package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type product struct {
	Id    int
	Name  string
	Price float32
}

func main() {

	//屬性必須大寫 表示公開 才可正確Marshal
	p1 := product{
		Id:    2,
		Name:  "Golang",
		Price: 10.56,
	}
	//物件轉Json
	jsonData, err := json.Marshal(p1)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	//Josn轉字串
	json := string(jsonData)
	fmt.Println(json)
	//Josn轉字串寫出到檔案
	ioutil.WriteFile("product.json", jsonData, 0666)

}
