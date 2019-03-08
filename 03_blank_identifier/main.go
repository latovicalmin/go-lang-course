package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	/* fmt.Println("test")
	res, _ := http.Get("http://www.klix.ba")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Println(page) */
	two()
}

func two() {
	res, err := http.Get("http://www.klix.ba")
	if err != nil {
		log.Fatal(err)
	}
	page, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", page)
}
