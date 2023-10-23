package main

import "fmt"

func main() {
    nums := []int{2,3,4}
    sum := 0

    for _, num := range nums {
        sum += num
    }

    fmt.Println("sum:", sum)

    for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }

    kvs := map[string]string{"a": "apple", "b": "banana"}

    for k,v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }

    for k := range kvs {
        fmt.Println("key:", k)
    }

    for i, c := range "go" {
        fmt.Println(i,c)
    }
}

//Range iterates over elements in a data structure
//Range in arrays and slices give index and the item itself
//Range in map iterates over key/value pairs. It can iterate over only keys too
//Range on strings iterates over Unicode points where first value is starting byte index of the rune and second is the rune itself
