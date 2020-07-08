package main

import (
    "fmt"
    "time"
)

func main() {

	fmt.Println("function start.")

    timer1 := time.NewTimer(2 * time.Second)
	fmt.Println("test1")
    <-timer1.C  // timer1.C just like a channel of timer1, so there should be value in this blocking channel, than the following code would be executed.
	fmt.Println("test2")
	fmt.Println("Timer 1 fired")

    timer2 := time.NewTimer(time.Second)
    go func() {
        <-timer2.C
        fmt.Println("Timer 2 fired")
	}()
	// a timer may be useful is that you can cancel the timer before it fires
    stop2 := timer2.Stop()
    if stop2 {
        fmt.Println("Timer 2 stopped")
	}
	
    time.Sleep(2 * time.Second)
}