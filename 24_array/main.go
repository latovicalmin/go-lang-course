package main

import (
	"fmt"
)

func main() {
	var arrayTest [10]int
	myFunc(arrayTest)

	var arrayChar [58]string

	for i := 65; i <= 122; i++ {
		arrayChar[i-65] = string(i)
	}

	fmt.Println(arrayChar)
}

func myFunc(arr [10]int) {
	arr[4] = 99
}
