package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
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
	uvi, _ :=
		os.OpenFile("uvi.json", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	rd := readPm25Json()
	defer rd.Close()
	b, _ := ioutil.ReadAll(rd)
	bfw := bufio.NewWriter(uvi)
	bfw.Write(b)
	bfw.Flush()
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
