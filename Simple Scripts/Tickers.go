package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(5 * time.Second)  // equals to start this ticker every 5 seconds

	done := make(chan bool, 1)

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("here is done")
				return  // return means the goroutine function end
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			// default:
			// 	fmt.Println("default")
			// 	return
			}
		}
	}()

	time.Sleep(16 * time.Second)
	ticker.Stop()
	done <- true
    fmt.Println("Ticker stopped")

}