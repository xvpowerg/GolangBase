package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func readCsv(csvr *csv.Reader) {
	//i := 0
	for {

		csvData, err := csvr.Read()
		if err == io.EOF {
			break
		}
		// if i == 0 {
		// 	i++
		// 	continue
		// }
		fmt.Printf("id:%s name:%s score1:%s score2:%s \n",
			csvData[0], csvData[1], csvData[2], csvData[3])
	}
}

func readAllCsv(csvr *csv.Reader) {
	csvDataArray, _ := csvr.ReadAll()
	for _, csvData := range csvDataArray {

		fmt.Printf("id:%s name:%s score1:%s score2 %s \n",
			csvData[0], csvData[1], csvData[2], csvData[3])
	}
}

func main() {

	f1, err := os.Open("student.csv")
	defer f1.Close()
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	csvr := csv.NewReader(f1)
	//readCsv(csvr)
	readAllCsv(csvr)

}
