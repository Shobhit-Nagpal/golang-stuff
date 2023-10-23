package main

import "fmt"

func zeroval(ival int) {
    ival = 0
}

func zeroptr(iptr *int) {
    *iptr = 0
}

func main() {
    i := 1
    fmt.Println("initial:", i)

    zeroval(i)
    fmt.Println("zeroval:", i)

    zeroptr(&i) //This gives the memory location of i, or a pointer to i
    fmt.Println("zeroptr:", i)

    fmt.Println("pointer:", &i)

}

// GO HAS POINTERS LET'S GOOOOOOOOOOOOOOO
//As we can see, zeroval takes in the copy of ival which is different from the one in the calling function
//zeroptr however takes in a *int as its parameter meaning it takes an int pointer. The *iptr dereferences the pointers from its memory location to the value in the address
//Assigning the value to a dereferenced pointer changes the value at referenced address
