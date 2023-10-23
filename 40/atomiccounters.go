package main

import (
    "fmt"
    "sync"
    "sync/atomic"
)

func main() {
    var ops atomic.Uint64

    var wg sync.WaitGroup

    for i := 0; i < 50; i++ {
        wg.Add(1)

        go func() {
            for c := 0; c < 1000; c++ {
                ops.Add(1)
            }

            wg.Done()
        }()
    }

    wg.Wait()

    fmt.Println("ops:", ops.Load())
    // We use Load method here as it is safe to automatically read a value even when other goroutines are updating it
    // If we did not use atomic integer and just did like ops++, we would get a different number because the gorountines would interfere with each other
    // We would also get data race failures when running the -race flag
}

// The primary mechanism for managing state in Go is communication over channels. There are other options as well like atomic counters which are accessible by goroutines
