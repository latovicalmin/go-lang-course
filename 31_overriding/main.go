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
	jmbg  string
}

func (p Person) Greeting() {
	fmt.Println("Hi")
}

func (s student) Greeting() {
	fmt.Println("Hey")
}

type human struct {
	first string
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
	fmt.Println(p1.Person.jmbg)
	fmt.Println(p1.jmbg)

	p1.Greeting()
	p1.Person.Greeting()

	h := &human{"Al", 23}
	fmt.Println(h)
	fmt.Println(h.first)
}
