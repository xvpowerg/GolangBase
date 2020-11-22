package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	path := `..\asset\pharmacy.csv`
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}
	r := csv.NewReader(f)
	_, _ = r.Read()
	pharmacy, err2 := r.Read()
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(pharmacy[0])
	fmt.Println(pharmacy[1])
	fmt.Println(pharmacy[2])
}
