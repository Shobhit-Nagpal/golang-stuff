package main

import (
    "fmt"
    "crypto/sha256"
)

func main() {
    s := "sha deeznuts"

    h := sha256.New()

    fmt.Println("slice of bytes for s:", []byte(s))
    fmt.Println("length slice of bytes for s:", len([]byte(s)))

    h.Write([]byte(s))

    bs := h.Sum(nil)

    fmt.Println(s)
    fmt.Printf("%x\n", bs)
}
