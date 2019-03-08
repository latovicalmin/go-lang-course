package main

import (
	"fmt"
)

func main() {
	student := make([]string, 35, 75)
	test := []string{}
	student[34] = "al"
	fmt.Println(student[0])
	fmt.Println(student == nil)
	fmt.Println(test == nil)

	fmt.Println("=====")
	arr := make([]int, 1)
	fmt.Println(arr[0])
	arr[0] = 7
	fmt.Println(arr[0])
	arr[0]++
	fmt.Println(arr[0])
	arr = append(arr, 4)
	fmt.Println(arr)
	fmt.Println(cap(arr))
	arr = append(arr, 9)
	fmt.Println(cap(arr))
}
