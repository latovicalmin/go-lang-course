package main

import (
	"fmt"
)

func myFunction(numbers []int, callback func(int)) {
	for _, n := range numbers {
		callback(n)
	}
}

func filter(numbers []int, callback func(int) bool) []int {
	var xs []int
	for _, val := range numbers {
		if callback(val) {
			xs = append(xs, val)
		}
	}
	return xs
}

func main() {
	myFunction([]int{1, 2, 3, 4, 7}, func(number int) {
		// fmt.Println(number)
	})

	xs := filter([]int{1, 2, 3, 4}, func(n int) bool {
		return n > 1
	})
	fmt.Println(xs)
}
