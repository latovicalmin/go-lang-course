package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			fmt.Println("Odd")
		} else {
			fmt.Println("Even")
		}
	}

	for index := 0; index < 10; index++ {
		for j := 0; j < index; j++ {
			fmt.Println(index, " - ", j)
		}

	}
}
