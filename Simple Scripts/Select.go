// Go’s select lets you wait on multiple channel operations. Combining goroutines and channels with select is a powerful feature of Go.
package main

import (
    "fmt"
    "time"
)

func main() {

    // c1 := make(chan string)
    // c2 := make(chan string)
	// // Note that the total execution time is only ~2 seconds since both the 1 and 2 second Sleeps execute concurrently.
	// // func like this is anoymous function, defined then execute
	// go func() {
    //     time.Sleep(1 * time.Second)
    //     c1 <- "one"
    // }()
    // go func() {
    //     time.Sleep(2 * time.Second)
    //     c2 <- "two"
    // }()

    // for i := 0; i < 2; i++ {  // means using select twice
    //     select {
    //     case msg1 := <-c1:
    //         fmt.Println("received", msg1)
    //     case msg2 := <-c2:
    //         fmt.Println("received", msg2)
    //     }
	// }
	

	// timeout
    c1 := make(chan string, 1)
    go func() {
        time.Sleep(2 * time.Second)
        c1 <- "result 1"
    }()

    select {
    case res := <-c1:  // the var res is only exists in this select block
        fmt.Println(res)
    case <-time.After(1 * time.Second):
        fmt.Println("timeout 1")
    }

    c2 := make(chan string, 1)
    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "result 2"
    }()
    select {
    case res := <-c2:
        fmt.Println(res)
    case <-time.After(3 * time.Second):
        fmt.Println("timeout 2")
	}

	// After execution done here, there would be a content in channel c1
	
	// for two channels, if c1 and c2 both have contents
	select {  // select just like switch-case, would only choose once case to execute
	case res1 := <- c1:  
		fmt.Println("select1, res1:  ", res1)
	case res2 := <- c2:
		fmt.Println("select1, res2:  ", res2)
	default:
		fmt.Println("default select")
	}

	fmt.Println("after pengzf testing")

	// fmt.Println(<- c1)  // c1 would be empty
	// fmt.Println(<- c2)

}