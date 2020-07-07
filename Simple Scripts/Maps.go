package main

import "fmt"

func main() {
	// still need to use make func to create map, make(map[key-type]val-type)
	m := make(map[string]int) 
	fmt.Println(m)

	n := map[string]int{"foo": 1, "bar": 2}  // another way to init map
	fmt.Println("map:", n)
	

	a, b := n["123"]  // a is the value, b is the bool, and if there is not key here, a  should be the default value of corresponding type of value
	fmt.Println(a)
	fmt.Println(b)
	


}