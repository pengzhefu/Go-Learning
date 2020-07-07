package main

import "fmt"

func zeroval(ival int) {  // in golang, the func can not change the value of input of the var if input the val
    ival = 0
}

func zeroptr(iptr *int) {  // can change the value of args when input the pointer
    *iptr = 0  // just like C, *pointer = var
}

func main() {
    i := 1
    fmt.Println("initial:", i)

    zeroval(i)
    fmt.Println("zeroval:", i)

    zeroptr(&i)
    fmt.Println("zeroptr:", i)

    fmt.Println("pointer:", &i)
}