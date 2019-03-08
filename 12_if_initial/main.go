package main

import (
	"fmt"
)

func main() {
	x := 4

	if test := x % 2; test == 0 {
		fmt.Println(test)
	}
}
