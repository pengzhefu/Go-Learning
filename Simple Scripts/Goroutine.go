package main

import (
    "fmt"
    "time"
)

func f(from string) {
    for i := 0; i < 3; i++ {
        fmt.Println(from, ":", i)
    }
}

func cal(a int , b int )  {
    c := a+b
    fmt.Printf("%d + %d = %d\n",a,b,c)
}

func main() {

    // f("direct")

	// // the following tow go blocks are two goroutines that run asynchronysouly
    // go f("goroutine")

    // go func(msg string) {
    //     fmt.Println(msg)
    // }("going")

    // time.Sleep(time.Second)
	// fmt.Println("done")
	
	for i :=0 ; i<10 ;i++{
        go cal(i,i+1)  //启动10个goroutine 来计算
    }
    time.Sleep(time.Second * 2)
}