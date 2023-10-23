package main

import (
    "fmt"
    "unicode/utf8"
)

func main() {
    const s = "สวัสดี"

    fmt.Println("Len:", len(s))

    for i := 0; i < len(s); i++ {
        fmt.Printf("%x ", s[i])
    }
    fmt.Println()

    fmt.Println("Rune count:",utf8.RuneCountInString(s))

    for idx, runeVal := range s {
        fmt.Printf("%#U starts at %d\n", runeVal, idx)
    }

    fmt.Println("\nUsing DecodeRuneInString")

    for i, w := 0, 0; i < len(s); i +=w {
        runeVal, width := utf8.DecodeRuneInString(s[i:])
        fmt.Printf("%#U starts at %d\n", runeVal, i)
        w = width

        examineRune(runeVal)
    }
}

func examineRune(r rune) {
    if r == 't' {
        fmt.Println("found tee")
    } else if r == 'ส' {
        fmt.Println("found so sua")
    }
}

//Strings in Go are read-only slice of bytes. The lang and the std library treat strings specially as containers of text encoded in UTF-8
//In other langs, strings are made of characters but in Go, the concept of a character is called rune
//Rune is an integer that represents a Unicode code point
//Values encoded in single quotes are rune literals
