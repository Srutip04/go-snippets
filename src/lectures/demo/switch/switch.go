package main

import "fmt"

func price() int {
	return 1
}

const (
	Economy    = 0
	Business   = 1
	FirstClass = 2
)

func main() {
	switch p := price(); {
	case p < 2:
		fmt.Println("cheap item")
	case p < 10:
		fmt.Println("Moderate")
	default:
		fmt.Println("Expensive")
	}

	ticket := Economy
	switch ticket{
	case Economy:
		fmt.Println("Economy Seating")
	case Business:
		fmt.Println("Business seating")
	case FirstClass:
		fmt.Println("FirstClass seating")
	default:
	     fmt.Println("other")
	}
}
