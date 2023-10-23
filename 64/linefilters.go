package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

//A line filter is a common type of program that reads input on stdin, processes it, and then prints some derived result to stdout. Ex: grep and sed

func main() {
        scanner := bufio.NewScanner(os.Stdin) //This gives us the Scan method

    for scanner.Scan() {
        ucl := strings.ToUpper(scanner.Text()) //Text gives us the current token 
        fmt.Println(ucl)
    }

    if err := scanner.Err(); err != nil {
        fmt.Println(os.Stderr, "error:", err)
        os.Exit(1)
    }
}
