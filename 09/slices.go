package main

import (
    "fmt"
    "slices"
)

func main() {
    var s []string
    fmt.Println("unint:", s, s == nil, len(s) == 0)

    s = make([]string, 3)
    fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s)
    fmt.Println("get:", s[2])

    fmt.Println("len:", len(s))

    s = append(s, "d")
    s = append(s, "e", "f")
    fmt.Println("apd:", s)

    c := make([]string, len(s))
    copy(c,s)
    fmt.Println("cpy:", c)

    l := s[2:5] //We get s[2], s[3], s[4]
    fmt.Println("sl1:", l)

    l = s[:5] // We get elements upto s[4]
    fmt.Println("sl2:", l)

    l = s[2:] //We start taking elements from s[2]
    fmt.Println("sl3:", l)

    t := []string{"g", "h", "i"}
    fmt.Println("dcl:", t)

    t2 := []string{"g", "h", "i"}
    if slices.Equal(t, t2) {
        fmt.Println("t == t2")
    }

    twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i+1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i +j
        }
    }
    fmt.Println("2d:", twoD)
}

//Slices are a datatype in Go which give a more powerful interface to sequence than arrays
//Slices are typed only by elements they contain, not number of elements like in arrays
//We use the built-in make to create a slice with non-zero length
//We need to accept the return value from append as we may get a new slice value
//Slices can be copied
//There is a "slice" operator with the syntax of slice[low:high]
//We can have multi-dimensional slices too ; Length of inner slices can vary unlike in multi-dimensional arrays
