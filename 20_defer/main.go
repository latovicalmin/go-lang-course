package main

import (
	"fmt"
)

func hello() {
	fmt.Println("Hello")
}

func world() {
	fmt.Println("World")
}

func main() {
	defer hello()
	world()
}
