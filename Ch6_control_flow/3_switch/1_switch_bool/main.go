package main

import "fmt"

func main() {

	type Action int
	const (
		PLAY Action = iota
		STOP
		EXIT
	)

	act := STOP
	switch act {
	case PLAY:
		fmt.Println("PLAY")
	case STOP:
		fmt.Println("STOP")
	case EXIT:
		fmt.Println("EXIT")
	}

	switch {
	case act == PLAY:
		fmt.Println("PLAY")
	case act == STOP:
		fmt.Println("STOP")
	case act == EXIT:
		fmt.Println("EXIT")
	}

}
