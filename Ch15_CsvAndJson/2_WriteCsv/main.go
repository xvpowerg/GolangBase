package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func writeCsv(csvw *csv.Writer) {
	var csvdata = [][]string{
		{"id", "name", "count"},
		{"1", "Iphone8", "10"},
		{"2", "Iphone11", "5"},
		{"3", "Iphone12", "50"},
	}
	for _, v := range csvdata {
		csvw.Write(v)
	}
	csvw.Flush()
}

func writeAllCsv(csvw *csv.Writer) {
	var csvdata = [][]string{
		{"id", "name", "count"},
		{"1", "Android O", "70"},
		{"2", "Android 11", "25"},
		{"3", "Golang", "50"},
	}
	csvw.WriteAll(csvdata)
	csvw.Flush()
}

func main() {
	f1, err := os.OpenFile("./test.csv",
		os.O_CREATE|os.O_WRONLY|os.O_TRUNC,
		0666)

	if err != nil {
		fmt.Println("err:", err)
	}
	defer f1.Close()

	csvw := csv.NewWriter(f1)
	//writeCsv(csvw)
	writeAllCsv(csvw)
}
