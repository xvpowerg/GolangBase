package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func writeAllCsv(csvw *csv.Writer) {
	var csvdata = [][]string{
		{"id", "name", "count"},
		{"1", "Android O", "70"},
		{"2", "Android 11", "25"},
		{"3", "Golang", "50"},
	}
	csvw.WriteAll(csvdata)
}

func main() {
	f1, err := os.OpenFile("test.csv",
		os.O_CREATE|os.O_WRONLY|os.O_TRUNC,
		0666)
	defer f1.Close()
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	csvw := csv.NewWriter(f1)
	writeAllCsv(csvw)
}
