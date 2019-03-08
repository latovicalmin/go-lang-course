package main

import (
	"fmt"
)

func main() {
	tempData := []float64{1, 2, 3, 4}
	fmt.Println(myFunc(tempData...))

	data := []float64{2, 4, 6, 8}
	n := average(data...)
	fmt.Println(n)
}

func myFunc(data ...float64) float64 {
	return data[2]
}

func average(sf ...float64) float64 {
	total := 0.0
	for _, v := range sf {
		total += v
	}

	return total / float64(len(sf))
}
