package main

import (
	"fmt"
)

type Square struct {
	side float64
}

func (s Square) area() float64 {
	return s.side * s.side
}

type Shape interface {
	area() float64
}

func info(z Shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

func test(v interface{}) {
	fmt.Println(v)
}

func PrintAll(vals []interface{}) {
	for _, val := range vals {
		fmt.Println(val)
	}
}

func main2() {
	s := Square{10}
	info(s)
	test(s)
}
