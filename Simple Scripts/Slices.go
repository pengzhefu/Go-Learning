package main

import "fmt"

func main() {

	s := make([]string, 3)  // using make func to create empty slice
	fmt.Println("init slice: ", s)

    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	s = append(s, "e", "f")  // using append 
	
	v := make([]string, 1)
	
	copy(v,s)  // copy(new, old), and the length depends on the new one
	fmt.Println("S: ", s)
	fmt.Println("v: ", v)  


	fmt.Println("slice: ", s[:3]) // support slice likes python, but the index can not larger than length

	// fmt.Println("index: ", s[-1])  // index could not be negative, unlike python

	t := []string{"g", "h", "i"}  // the way like 
    fmt.Println("dcl:", t)

}