package main

import "fmt"

func main() {
    messages := make(chan string) //This is how we create a channel

    go func() {messages <- "ping"}() // Send a value into the channel using <-

    msg := <-messages // Here we receive the value from the channel
    fmt.Println(msg)
}

// Channels are pipes that connect concurrent gorountines
// We can send values into channels from one goroutine and receive those values from another goroutine
