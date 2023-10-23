package main

import (
    "fmt"
    "sync"
)

type Container struct {
    mu sync.Mutex
    counters map[string]int
}

func (c *Container) inc(name string) {
    c.mu.Lock() //Lock mutex before accessing counters
    defer c.mu.Unlock() //Unlock mutex using defer ; defer will run this after this function returns the value
    c.counters[name]++
}

func main() {
    c := Container{
        counters: map[string]int{"a": 0, "b": 0},
    }

    var wg sync.WaitGroup

    doIncrement := func(name string, n int) {
        for i := 0 ; i < n ; i++ {
            c.inc(name)
        }
        wg.Done()
    }

    wg.Add(3)
    go doIncrement("a", 10000)
    go doIncrement("a", 10000)
    go doIncrement("b", 10000)

    wg.Wait()
    fmt.Println(c.counters)
}


// We can use mutexes to safely access data across multiple gorountines in complex state
// In this example, for Container, since we want to update counters from multiple gorountines, we use mutex to synchronize access
// Note: Mutexes must not be copied. If this struct is passed around, it should be done by pointers 

