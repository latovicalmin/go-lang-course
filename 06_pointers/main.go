package main

import (
	"fmt"
)

func main() {
	a := 42

	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a

	fmt.Println(b)
	fmt.Println(&b)
	fmt.Println(*b)

	changeNum(a)
	fmt.Println(a)

	changenNum2(&a)
	fmt.Println(a)
}

func changeNum(num int) {
	num += 3
}

func changenNum2(num *int) {
	*num += 5
}

func changeName(val) {
	val.name = ""
}
