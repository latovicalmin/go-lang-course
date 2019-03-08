package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func (p person) fullName() string {
	return p.first + " " + p.last
}

func (p person) changeName(name string) {
	p.first = name
}

func (p *person) changeName2(name string) {
	p.first = name
}

func main() {
	p1 := person{"Al", "La", 23}
	p2 := person{"Ha", "De", 23}
	p1.changeName("AA")  // This will not change
	p1.changeName2("AA") // This will change
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.fullName())
}
