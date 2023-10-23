package main

import "fmt"

func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

func main() {
    nextInt := intSeq()

    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())

    newInts := intSeq()
    fmt.Println(newInts())
}

//Go supports anonymous functions which can form closures
//Anon functions are useful when you want to define a function inline without having to name it

// The intSeq() function returns a function (which itself returns an int). The function that's returned forms a closure over i. So when intSeq() is called by nextInt, the i is initialized and when we call nextInt(), the anonymous function gets to work where we alredy have i and it increments i and gives it back 
