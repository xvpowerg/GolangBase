package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

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
	var objMap map[string]interface{}
	//注意必須objMap是point
	json.Unmarshal(jsonData, &objMap)

	fmt.Println(objMap["Id"].(float64))
	fmt.Println(objMap["Name"].(string))
	fmt.Println(objMap["Price"].(float64))

}
