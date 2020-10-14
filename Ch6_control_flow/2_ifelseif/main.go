package main

import "fmt"

func main() {
	//紫外線指數 https://www.cwb.gov.tw/V8/C/W/MFC_UVI_Map.html
	uva := 6

	if uva >= 0 && uva <= 2 {
		fmt.Println("低量級")
	} else if uva >= 3 && uva <= 5 {
		fmt.Println("中量級")
	} else if uva >= 6 && uva <= 7 {
		fmt.Println("高量級")
	} else if uva >= 8 && uva <= 1 {
		fmt.Println("過量級")
	} else {
		fmt.Println("危險級")
	}

}
