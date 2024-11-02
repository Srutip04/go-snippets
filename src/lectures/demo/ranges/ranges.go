package main

import (
	"fmt"
)

func main() {
	slice := []string{"Hello","World","!"}

	for i, element := range slice{
		fmt.Println(i,element,":")

		for _, ch := range element{
			fmt.Println("   %q\n", ch)
		}
	}
}
