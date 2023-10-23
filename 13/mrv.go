package main

import "fmt"

func vals() (int, int) {
    return 3,7
}

func main() {
    a, b := vals()
    fmt.Println(a)
    fmt.Println(b)

    _,c := vals() //Using blank identifier for the value we don't want
    fmt.Println(c)
}

//Go can return multiple values from a function
//This feature is usually used in idiomatic Go where we return the result and the error values
