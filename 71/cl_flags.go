package main

import (
    "flag"
    "fmt"
)

func main() {
    wordPtr := flag.String("word", "foo", "a string") //flag.String returns a string pointer and NOT a string value

    numPtr := flag.Int("numb", 42, "an int")
    forkPtr := flag.Bool("fork", false, "a bool")

    var svar string
    flag.StringVar(&svar, "svar", "bar", "a string var") //Note that we're passing a pointer here

    flag.Parse() //Once all flags are declared, use Parse to execute command line parsing

    fmt.Println("word:", *wordPtr)
    fmt.Println("numb:", *numPtr)
    fmt.Println("fork:", *forkPtr)
    fmt.Println("svar", svar)
    fmt.Println("tail:", flag.Args())
}
