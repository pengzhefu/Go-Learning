package main

import "fmt"

type rect struct {
    width, height int
}

// the following two funcs are the methods of struct rect, both pointer and struct itself could be ok
func (r *rect) area() int {
    return r.width * r.height
}

func (r rect) perim() int {  // similar as func, but the name is after args and before return type
    return 2*r.width + 2*r.height
}

func main() {
    r := rect{width: 10, height: 5}

    fmt.Println("area: ", r.area())
    fmt.Println("perim:", r.perim())

    rp := &r
    fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
	
	fmt.Println("after change the value by struct itself")
	r.height = 100
	fmt.Println("area: ", r.area())
    fmt.Println("perim:", r.perim())

    fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}