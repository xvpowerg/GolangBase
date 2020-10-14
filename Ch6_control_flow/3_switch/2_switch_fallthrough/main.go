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
		fallthrough
	case EXIT:
		fmt.Println("EXIT")
	}

}
