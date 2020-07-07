//Interfaces are named collections of method signatures.

package main

import (
    "fmt"
    "math"
)

type geometry interface {  // two methods for "geometry" interface
    area() float64
    perim() float64
}

type rect struct {
    width, height float64
}
type circle struct {
    radius float64
}

type bilibala struct {
	params float64
}

func (r rect) area() float64 {
    return r.width * r.height
}
func (r rect) perim() float64 {
    return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}

func (b bilibala) area() float64 {
	return b.params
}

func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

func main() {
    r := rect{width: 3, height: 4}
    c := circle{radius: 5}
	b := bilibala{12.1}

    measure(r)
	measure(c)
	// measure(b)  // need to implement all the method of interfaces so that use the function of interface
}