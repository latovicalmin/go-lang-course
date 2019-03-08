package main

import (
	"fmt"
)

const x string = "won't be changed"
const y = 42

const (
	_ = iota
	a
	b
	c
	d = iota * 10
)

func main() {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(c)
	fmt.Println(d)
}
