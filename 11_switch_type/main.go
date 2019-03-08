package main

import (
	"fmt"
)

func main() {
	type Contact struct {
		greeting string
		name     string
	}

	checkType(4)
	checkType(true)
	checkType("gfdgfgd")
}

func checkType(x interface{}) {
	switch x.(type) {
	case string:
		fmt.Println("Stringg")
	case int:
		fmt.Println("Intt")
	case bool:
		fmt.Println("Booll")
	}
}
