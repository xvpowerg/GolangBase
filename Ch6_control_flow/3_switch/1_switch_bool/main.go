package main

import "fmt"

func main() {

	name := "Ken"
	switch name {
	case "Vivin":
		fmt.Println("業務")
	case "Ken":
		fmt.Println("經理")
	case "Lucy":
		fmt.Println("PM")
	}

	switch {
	case name == "Vivin":
		fmt.Println("業務")
	case name == "Ken":
		fmt.Println("經理")
	case name == "Lucy":
		fmt.Println("PM")

	}

	//可使用常數
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
