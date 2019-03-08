package main

import (
	"fmt"
)

func main() {
	firstSlice := []string{"4a", "7", "9"}
	secondSlice := []string{"aa", "bbb"}
	thirdSlice := append(firstSlice, secondSlice...)
	fmt.Println(thirdSlice)
	thirdSlice = append(thirdSlice[:2], thirdSlice[3:]...)
	fmt.Println(thirdSlice)

	greeting := make([]string, 3, 5)
	greeting = append(greeting, "hello")
	fmt.Println(greeting)
	fmt.Println(len(greeting))
	fmt.Println(cap(greeting))

	mySlice := []int{1, 5, 7, 8}

	fmt.Println(len(mySlice))
	fmt.Println(mySlice)

	s := make([]string, 7)
	s[4] = "lol"
	fmt.Println("slices: ", s)

	s = append(s, "bro", "hello")
	fmt.Println("s: ", s)

	c := make([]string, len(s))
	fmt.Println("cap: ", cap(c))
	fmt.Println("len: ", len(c))
	copy(c, s)

	s[4] = "sta"

	fmt.Println("c: ", c)
	fmt.Println("s: ", s)

	l := s[2:5]
	fmt.Println("l: ", l)

	l = s[5:]
	fmt.Println("l: ", l)

	newSlice := []string{"a", "b", "c", "g", "m", "z"}
	fmt.Println(len(newSlice))
	fmt.Println(newSlice[2:4])
	fmt.Println(string("myString"[2]))

	myFunc(newSlice)
	fmt.Println(string(newSlice[2]))

	d := make([]int, 2, 4)
	fmt.Println("d: ", d)
	fmt.Println("d len: ", len(d))
	fmt.Println("d cap: ", cap(d))
}

func myFunc(arr []string) {
	arr[2] = "H"
}
