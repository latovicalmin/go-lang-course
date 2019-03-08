package main

import (
	"fmt"
)

func main() {
	fmt.Println([]byte("Hello"))

	for index := 0; index < 10; index++ {
		for j := 0; j < index; j++ {
			// fmt.Println(index, " - ", j)
		}
	}

	x := 2

	for x < 10 {
		if x == 7 {
			break
		}
		//fmt.Println(x)
		x++
	}

	y := 0

	for y <= 20 {
		y++
		if y%2 == 0 {
			continue
		}

		fmt.Println(y)
	}

	for i := 5000; i < 5100; i++ {
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
	}
}
