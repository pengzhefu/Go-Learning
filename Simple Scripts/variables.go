package main

import "fmt"

func main() {
	var a = "zfpeng"

	fmt.Println(a)

	var b string = "another zfpeng"
	fmt.Println(b)

	var c int = 123
	fmt.Println(c)

	// var c,d = 456, 789 // not allowed to redeclare this var, not like in scala var
	// fmt.Println(c)
	var d, f  = 7, 8 
	fmt.Println(d)
	fmt.Println(f)

	g := 8
	fmt.Println(g)

	var e int
	fmt.Println(e)

	var s string
	fmt.Println(s)

	g += 1  // var could not be redeclared, but the value could be changed.
	fmt.Println("g after add himself: ", g)
}