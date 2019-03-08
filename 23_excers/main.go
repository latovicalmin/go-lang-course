// Multiple returns
// Func expression
// Variadic func
// Bool
// Variadic again
package main

import (
	"fmt"
)

func main() {
	testArray := []int{2, 3, 4}
	res1, res2 := multiRes(testArray)
	testArray2 := []float64{2, 5, 3}
	fmt.Println(res1, "---", res2)
	fmt.Println(varFun(testArray2...))

	// Func expression
	testFun := func(par []int) (int, bool) {
		return par[1], par[1]%2 == 0
	}
	r1, r2 := testFun(testArray)
	fmt.Println(r1, r2)
}

// Multiple returns
func multiRes(par []int) (int, int) {
	return par[0], par[2]
}

// Variadic
func varFun(par ...float64) float64 {
	return par[1] / 2
}
