package main

import "fmt"

func ping(pings chan<- string, msg string) {
	fmt.Println("ping")
	pings <- msg
	// fmt.Println("ping")
}

func pong(pings <-chan string, pongs chan<- string) {
	fmt.Println("pong")
	msg := <-pings
	pongs <- msg
	// fmt.Println("pong")
}

func main() {
    pings := make(chan string, 1)
    pongs := make(chan string, 1)
    go ping(pings, "passed message")
	go pong(pings, pongs)
	// <-pongs  // here to test what we said before in Channel2confirmGoroutine
    fmt.Println(<-pongs)
}