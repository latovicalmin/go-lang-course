package main

import (
	"fmt"
)

func main() {
	a := "43"
	fmt.Println("a - ", a, " yo")
	fmt.Println("a's address - ", &a)
	input()
}

func input() {
	var x float64
	fmt.Scan(&x)
	fmt.Println("ola: ", x+3)
}
