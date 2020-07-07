package main

import "fmt"

func main() {
	
	// arrays in go just like arrays in java
    var a [5]int
    fmt.Println("emp:", a)

    a[4] = 100  // differ with declared to var, no need to use :
    fmt.Println("set:", a)
    fmt.Println("get:", a[4])

    fmt.Println("len:", len(a))

    b := [5]int{1, 2, 3, 4, 5}  // declare of array with 
    fmt.Println("dcl:", b)

    var twoD [2][3]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}