package main

import (
    "fmt"
    "time"
)

// this example is to show that use channel to wait goroutine finish

func worker(done chan bool) {
    fmt.Print("working...")
    time.Sleep(time.Second)
    fmt.Println("done")

    done <- true
}

func main() {

    done := make(chan bool, 1)
    go worker(done)

	<-done  // If you removed the <- done line from this program, the program would exit before the worker even started.
	// the reason is that program need to get the output from the channel done, but if program
	// cannot get the output, it will just stop here without exiting and wait for the complete
	// of goroutine
}