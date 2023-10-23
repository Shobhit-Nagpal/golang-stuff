package main

import (
    "fmt"
    "time"
)

func worker(done chan bool) {
    fmt.Println("Working...")
    time.Sleep(time.Second)
    fmt.Println("Done...")

    done <- true // Send a value to notify that we're done
}

func main() {
    done := make(chan bool, 1)
    go worker(done) //Start a worker gorountine and give the channel for it to notify

    <-done // Block until we receive a notification from the worker on channel. If we removed this line, the program would exit before the worker even started working
}


//We can use channels to synchronize execution across goroutines
// When waiting for multiple gorountines to finish, we may prefer a WaitGroup

