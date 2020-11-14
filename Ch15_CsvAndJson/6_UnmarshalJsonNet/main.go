package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const minReadBufferSize = 16
const maxConsecutiveEmptyReads = 100
const (
	url string = `https://data.epa.gov.tw/api/v1/uv_s_01?limit=1000&api_key=9be7b239-557b-4c10-9775-78cadfc555e9&format=json`
)

func readPm25Json() io.ReadCloser {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("err:", err)
	}
	re := resp.Body
	return re
}

func downloadJson() {

	rd := readPm25Json()
	defer rd.Close()
	b, _ := ioutil.ReadAll(rd)
	ioutil.WriteFile("uvi.json", b, 0666)
}

type UviInfo struct {
	County      string
	PublishTime string
	SiteName    string
	UVI         string
}

type Uvi struct {
	Records []UviInfo
}

func JsonToObj() {
	rd := readPm25Json()
	defer rd.Close()
	b, _ := ioutil.ReadAll(rd)
	var uvi Uvi
	json.Unmarshal(b, &uvi)
	fmt.Println(uvi)
}

func main() {
	downloadJson()
	//JsonToObj()
}
