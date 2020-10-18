package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type product struct {
	Id    int
	Name  string
	Price float32
}

func main() {
	f1, err := os.OpenFile("./product.json",
		os.O_CREATE|os.O_WRONLY|os.O_TRUNC,
		0666)
	fw := bufio.NewWriter(f1)

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
	}
	//Josn轉字串
	json := string(jsonData)
	fmt.Println(json)
	//Josn轉字串寫出到檔案
	fw.WriteString(json)
	fw.Flush()
}
