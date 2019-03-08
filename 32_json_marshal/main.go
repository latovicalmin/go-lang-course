package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string
	Last  string
	age   int
}

func main() {
	p1 := Person{"Al", "La", 23}
	bs, _ := json.Marshal(p1)
	fmt.Println(bs)
	fmt.Printf("%T \n", bs)
	fmt.Println(string(bs))

	fmt.Println("-------------")
	bsNew := []byte(`{"First": "James", "Last": "Bond", "Age": 20}`)
	var p2 Person
	json.Unmarshal(bsNew, &p2)
	fmt.Println(p2.First)
	fmt.Println(p2)
}
