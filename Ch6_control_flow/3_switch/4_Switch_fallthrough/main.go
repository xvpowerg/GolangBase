package main

import "fmt"

func main() {

	type Action int
	const (
		PLAY Action = iota
		STOP
		PAUSE
	)

	act := PLAY
	switch act {
	case PLAY:
		fmt.Println("PLAY")
		fallthrough
	case STOP:
		fmt.Println("STOP")
	case PAUSE:
		fmt.Println("PAUSE")
	}

}
