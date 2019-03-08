package main

import (
	"fmt"
)

func sayAllo() func() string {
	return func() string {
		return "Alloo"
	}
}

func main() {
	myFunc := func() {
		fmt.Println("Hello")
	}

	myFunc()

	newFunc := sayAllo()
	fmt.Printf("%T\n", newFunc)

	fmt.Println(newFunc())
}
