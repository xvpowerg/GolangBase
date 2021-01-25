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

type product2 struct {
	Id    int     `json:"myId"`
	Name  string  `json:"myName"`
	Price float32 `json:"myPrice"`
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
	jsonStr := string(jsonData)
	fmt.Println(jsonStr)
	//Josn轉字串寫出到檔案
	ioutil.WriteFile("product.json", jsonData, 0666)

	//如果希望輸出的json可自訂屬性名稱 可使用如product2 方式加上`json:"轉換後的名稱"`
	p2 := product2{
		Id:    5,
		Name:  "JavaScript",
		Price: 956,
	}
	jsonData2, _ := json.Marshal(p2)
	fmt.Printf(string(jsonData2))

}
