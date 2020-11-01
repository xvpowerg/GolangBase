package main

import "fmt"

func main() {

	//可使用常數模擬Enum
	type Action int
	const (
		PLAY Action = iota
		STOP
		PAUSE
	)
	type Fruit int
	const (
		Banana Fruit = iota
		Apple
		Kiwi
	)

	var act Action = PAUSE
	switch act {
	case PLAY:
		fmt.Println("PLAY")
	case STOP:
		fmt.Println("STOP")
	case PAUSE:
		fmt.Println("PAUSE")
	}

}
