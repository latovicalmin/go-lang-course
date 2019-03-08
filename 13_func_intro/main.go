package main

import (
	"fmt"
)

func main() {
	greeting(getFname(), "lat")
}

func greeting(fname, lname string) {
	fmt.Println(fname, lname)
}

func getFname() string {
	return "alm"
}
