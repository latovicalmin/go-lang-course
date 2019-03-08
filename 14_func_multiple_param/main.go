package main

import (
	"fmt"
)

func main() {
	x := 4

	fmt.Println(calcualte(x))
}

func calcualte(x int) (int, int) {
	return x + 2, x + 3
}
