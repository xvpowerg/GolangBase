package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func readCsv(csvr *csv.Reader) {
	for {
		csvData, err := csvr.Read()
		if err == io.EOF {
			break
		}
		fmt.Printf("id:%s name:%s score1:%s score2 %s \n",
			csvData[0], csvData[1], csvData[2], csvData[3])
	}
}

func readAllCsv(csvr *csv.Reader) {
	csvDatas, _ := csvr.ReadAll()
	for _, csvData := range csvDatas {
		fmt.Printf("id:%s name:%s score1:%s score2 %s \n",
			csvData[0], csvData[1], csvData[2], csvData[3])
	}
}

func main() {

	f1, err := os.Open("student.csv")
	if err != nil {
		fmt.Println("err:", err)
	}
	defer f1.Close()

	csvr := csv.NewReader(f1)
	readAllCsv(csvr)

}
