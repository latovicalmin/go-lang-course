package main

import (
	"fmt"

	"github.com/latovic/course/02_variables/testpack"
)

var bi int = 4

func wrapper() func() int {
	return func() int {
		bi++
		return bi
	}
}

func main() {
	var a bool
	var x, y string = "stored in x", "stored in y"
	var b = 24
	c := 44
	fmt.Println("%v ", a)
	fmt.Println("%v ", b)
	fmt.Println("%T", c)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(c)

	lolFunc := func() string {
		return "wazza"
	}

	testpack.PrintMyName()
	fmt.Println(lolFunc())

	fmt.Println(bi)
	myFunc := wrapper()
	myFunc()
	fmt.Println(bi)

}
