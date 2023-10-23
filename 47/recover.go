package main

import "fmt"

func mayPanic() {
    panic("a problem")
}

func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered. Error:\n", r)
        }
    }()

    mayPanic()

    fmt.Println("After mayPanic()")
}


//You can recover from a panic in Go. A recover can stop a panic from aborting the program and let it continue with the execution. For example, if you were using http server

//Recover must be called within a deferred function. When the enclosing function panics, the defer will activate and a recover call within it will catch the panic

//Return value of recover is the error raised in the call to panic
