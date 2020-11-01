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

}
