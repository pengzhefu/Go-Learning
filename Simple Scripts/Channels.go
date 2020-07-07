package main

import "fmt"

func main() {

    messages := make(chan string)  // using make to create a channel, key word: chan, and this is the ordinary channel

    go func() { messages <- "ping" }() // Send a value into a channel using the channel <- syntax.

    msg := <-messages  // The <-channel syntax receives a value from the channel
    fmt.Println(msg)
}