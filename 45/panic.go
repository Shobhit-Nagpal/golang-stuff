package main

import "os"

func main() {
    panic("a problem")

    _, err := os.Create("/tmp/file")
    if err != nil {
        panic(err)
    }
}

// Panic means something went unexpectedly wrong. Mostly, we use it to fail fast on errors that should not occur during normal op or we aren't perpared to handle gracefully\
// When first panic fires, the program exits without reaching rest of the code
