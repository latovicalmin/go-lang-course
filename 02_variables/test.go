package main

import "fmt"

func main() {
	sumNum := sumNum(6)
	fmt.Println(sumNum) // bad practice
}

func sumNum(number int) int {
	fmt.Println(testVar)
	return 33 + number
}

var testVar int = 5
