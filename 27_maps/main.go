package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println(m)

	delete(m, "k2")

	v, ok := m["k1"]
	fmt.Println(v, ok)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(n)

	for key, value := range n {
		fmt.Println("Index:", key, "value:", value)
	}
}
