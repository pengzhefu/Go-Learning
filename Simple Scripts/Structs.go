package main

import "fmt"

type person struct {
    name string
    age  int
}
// You can safely return a pointer to local variable as a local variable will survive the scope of the function.
func newPerson(name string) *person {  // the return type is pointer

    p := person{name: name}
    p.age = 42
    return &p
}

func main() {

	// the following four methods just to create the struct
    fmt.Println(person{"Bob", 20})

    fmt.Println(person{name: "Alice", age: 30})

    fmt.Println(person{name: "Fred"})

    fmt.Println(&person{name: "Ann", age: 40})  // it should be the address

    fmt.Println(newPerson("Jon"))  // it should be address

    s := person{name: "Sean", age: 50}
    fmt.Println(s.name)

    sp := &s
    fmt.Println(sp.age)  // get the val of struct can use both the struct itself and the pointer of struct

    sp.age = 51
	fmt.Println(sp.age)

	
	newP := newPerson("tb")  // newP is a pointer
	fmt.Println(newP)
	fmt.Println(*newP)  // get the val over the pointer

	// the following two sentences demonstrates that both using pointer and struct itself could change the value of struct
	newP.age = 100
	fmt.Println(*newP)

	(*newP).age = 200
	fmt.Println(*newP)
}