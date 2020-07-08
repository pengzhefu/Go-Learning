package main

import "fmt"

func main() {
    jobs := make(chan int, 5)
    done := make(chan bool)

    go func() {
        for {  // here to mention how to for loop the channel
            j, more := <-jobs
            if more {
                fmt.Println("received job", j)
            } else {
                fmt.Println("received all jobs: ", more)
                done <- true
                return
            }
        }
    }()

    for j := 1; j <= 3; j++ {
        jobs <- j
        fmt.Println("sent job", j)
    }
    close(jobs)  // when channel closes, the content in the channel still could be read
    fmt.Println("sent all jobs")

    <-done  // here is the blocking channel, so should wait till there is content in this channel
}