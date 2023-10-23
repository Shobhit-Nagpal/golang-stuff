package main

import (
    "fmt"
    "os"
)

type point struct {
    x, y int
}

func main() {
    p := point{1,2}
    fmt.Printf("struct1: %v\n", p) //Prints instance of our point struct

    fmt.Printf("struct2: %+v\n", p) //Includes field names

    fmt.Printf("struct3: %#v\n", p) //Prints Go syntax representation of value

    fmt.Printf("type: %T\n", p) //Prints type

    fmt.Printf("bool:  %t\n", true) //For bool

    fmt.Printf("int:  %d\n", 123) //%d is standard, base-10 formatting

    fmt.Printf("bin:  %b\n", 14) //Binary rep

    fmt.Printf("char:  %c\n", 33) //Prints character corresponding to integer

    fmt.Printf("hex:  %x\n", 456) // %x provides hex encoding

    fmt.Printf("float1:  %f\n", 78.9) 

    fmt.Printf("float2:  %e\n", 123400000.0) //%e amd %E use float in scientific notation
    fmt.Printf("float3:  %E\n", 123400000.0)

    fmt.Printf("str1: %s\n", "\"string\"")

    fmt.Printf("str2: %q\n", "\"string\"") //To double quote strings as in Go source

    fmt.Printf("str3: %x\n", "hex this") //Renders string in base-16

    fmt.Printf("pointer: %p\n", &p) //Pointer

    fmt.Printf("width1: |%6d|%6d|\n", 12, 345) //To specify use number after the % in the verb

    fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)

    fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45) //The - is for left-justify

    fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")

    fmt.Printf("width4: |%-6s|%-6s|\n", "foo", "b")

    s := fmt.Sprintf("sprintf: a %s", "string") //Formats and returns string without printing
    fmt.Println(s)

    fmt.Fprintf(os.Stderr, "io: an %s\n", "error") //We can format + print to io.Writers other than os.Stdout using Fprintf
}
