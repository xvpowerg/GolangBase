package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

//注意屬性必須是大寫開頭
type product struct {
	Id    int
	Name  string
	Price float32
}

func main() {
	f1, err := os.Open("./product.json")
	if err != nil {
		fmt.Println("err:", err)
	}
	r := bufio.NewReader(f1)
	scan := bufio.NewScanner(r)
	jsonData := make([]byte, 0)
	for scan.Scan() {
		jsonData = append(jsonData, scan.Bytes()...)
	}
	var p product
	//注意必須是point
	json.Unmarshal(jsonData, &p)
	fmt.Println(p)

}
