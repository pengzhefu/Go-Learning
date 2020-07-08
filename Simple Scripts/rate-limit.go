package main

import (
    "fmt"
    "time"
)

func main() {

    // requests := make(chan int, 5)
    // for i := 1; i <= 5; i++ {
    //     requests <- i
    // }
    // close(requests)

    // limiter := time.Tick(200 * time.Millisecond)

    // for req := range requests {
    //     <-limiter
    //     fmt.Println("request", req, time.Now())
	// }
	
	// the following code is for short bursts of requests

	// For the second batch of requests we serve the first 3 immediately 
	// because of the burstable rate limiting, then serve the remaining 2 with ~200ms delays each.

    burstyLimiter := make(chan time.Time, 3)

	// here to decide how many goroutines execute simutaneously for the begining of burstlimit
    for i := 0; i < 3; i++ {
        burstyLimiter <- time.Now()
    }

	// here to define how long should be the time interval
    go func() {
        for t := range time.Tick(10 * time.Second) {
            burstyLimiter <- t
		}
    }()

    burstyRequests := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        burstyRequests <- i
    }
    close(burstyRequests)
    for req := range burstyRequests {
		fmt.Println("wait for channel")
        <-burstyLimiter
        fmt.Println("request", req, time.Now())
    }
}