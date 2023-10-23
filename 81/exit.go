package main

import (
    "fmt"
    "os"
)

func main() {
    defer fmt.Println("!")
    os.Exit(3)
}

//defers will not be run when using os.Exit
