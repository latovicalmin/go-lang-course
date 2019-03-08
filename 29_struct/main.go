package main

import (
	"fmt"
)

type foo int

type Person struct {
	jmbg  string
	photo string
}

type student struct {
	Person
	first string "test"
	last  string
	age   int
}

func main() {
	p1 := student{
		Person: Person{
			jmbg:  "05",
			photo: "3213131",
		},
		last: "Al", first: "La", age: 23,
	}
	fmt.Println(p1)
	fmt.Println(p1.jmbg)

	var br foo = 4
	var nmb = 5

	fmt.Println(int(br) + nmb)
}
