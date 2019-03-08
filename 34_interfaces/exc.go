package main

import (
	"fmt"
	"sort"
)

type person struct {
	First string
	Age   int
}

type myint interface{}

type people []string

func (v people) Len() int           { return len(v) }
func (v people) Swap(i, j int)      { v[i], v[j] = v[j], v[i] }
func (v people) Less(i, j int) bool { return v[i] < v[j] }
func (v person) String() string {
	return "test " + v.First + " age: " + string(v.Age)
}

func main() {
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}
	s := []string{"Zeno", "John", "Al", "Jenny"}
	n := []int{7, 4, 7, 2, 9, 19, 12, 32, 3}

	test := person{"alm", 23}
	fmt.Println(test)
	vals := []myint{studyGroup, test, 4, s, n}
	fmt.Println(vals)

	sortPeople(studyGroup)
	stringSort(s)
	intSort(n)
	fmt.Println(studyGroup)
	fmt.Println(s)
	fmt.Println(n)

	fmt.Println("======")
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	sort.Sort(sort.Reverse(sort.IntSlice(n)))

	fmt.Println(studyGroup)
	fmt.Println(s)
	fmt.Println(n)

	fmt.Println("======")
	fmt.Println(sort.IsSorted(studyGroup))

	persons := []person{{"Al", 23}, {"Ha", 22}, {"Bi", 25}}
	fmt.Println(persons)

}

func sortPeople(v people) {
	sort.Sort(people(v))
}

func stringSort(v []string) {
	sort.Strings(v)
}

func intSort(v []int) {
	sort.Ints(v)
}
