package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	jsonData, _ := ioutil.ReadFile("product.json")
	var objMap map[string]interface{}
	//注意必須objMap是point
	json.Unmarshal(jsonData, &objMap)
	fmt.Println(objMap["Id"].(float64))
	fmt.Println(objMap["Name"].(string))
	fmt.Println(objMap["Price"].(float64))

}
