package main

import (
    "fmt"
    "time"
)

func main() {
    timer1 := time.NewTimer(2 * time.Second)

    <-timer1.C
    fmt.Println("Timer 1 fired")

    timer2 := time.NewTimer(time.Second)

    go func() {
        <- timer2.C
        fmt.Println("Timer 2 fired")
    }()

    stop2 := timer2.Stop()

    if stop2 {
        fmt.Println("Timer 2 has stopped")
    }

    time.Sleep(2 * time.Second)
}


//Timers represent a single event in the future
//Tell timer how long to wait, it provides a channel that will be notified at the time
