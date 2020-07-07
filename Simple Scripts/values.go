package main

import "fmt"

func main() {
	fmt.Println("here is the " + "start:")

	fmt.Println("2+2 = ", 4)
	// fmt.Println("2+2 = " +  4), wrong, the concat of int and string is like in python

	fmt.Println("7.0 / 3.0 = ", 7.0 / 3.0)

	fmt.Println("7 / 3 = ", 7/3)

	fmt.Println(true && false)

	fmt.Println("test of bool: ", true || false)
}