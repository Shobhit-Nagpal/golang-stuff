package main

import "fmt"

func main() {
    messages := make(chan string, 2)

    messages <- "buffered"
    messages <- "channel"

    fmt.Println(<-messages)
    fmt.Println(<-messages)
}

// By default, channels are unbuffered meaning that they will only accept sends if there is a corresponding receives to receive the sent value
// Buffered channels accept a limited amount of values without a corresponding receiver for those values
