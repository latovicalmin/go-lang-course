package main

import (
	"fmt"
)

func main() {
	x := 1

	switch x {
	case 1:
		fmt.Println("One")
		fallthrough
	case 5:
		fmt.Println("Five")
	default:
		fmt.Println("Default")
	}
}
