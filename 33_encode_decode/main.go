package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	First string
	Last  string
	age   int
}

func main() {
	p1 := Person{"Al", "La", 23}
	json.NewEncoder(os.Stdout).Encode(p1)

	var p2 Person
	rdr := strings.NewReader(`{"First": "Jame", "Last": "Bondic"}`)
	json.NewDecoder(rdr).Decode(&p2)
	fmt.Println(p2)
}
