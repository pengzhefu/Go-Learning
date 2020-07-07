package main

import (
	"fmt"
	"math"
)

func main() {

	const a = "here is a constant var"
	fmt.Println(a)

	const b float64 = 7.0 / 3.0
	fmt.Println(b)

	const c int = 7 / 3
	fmt.Println(c)

	fmt.Println(math.Sin(b))

	// b += 1  //the value of constants could not be changed.

}