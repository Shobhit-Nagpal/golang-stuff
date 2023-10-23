package main

import (
    "fmt"
    "bytes"
    "regexp"
)

func main() {
    match, _ := regexp.MatchString("p([a-z]+)ch", "peach") //Pattern matches string or not
    fmt.Println(match)

    r, _ := regexp.Compile("p([a-z]+)ch")  //For other regex tasks, we can Compile an optimized Regex struct 

    // Many methods that are available on these structs
    
    fmt.Println(r.MatchString("peach")) //Same match test

    fmt.Println(r.FindString("peach punch")) // Finds the match for regexp

    fmt.Println("idx:", r.FindStringIndex("peach punch")) //Finds first match but returns the start and end indexes for the match instead of matching text
    
    fmt.Println(r.FindStringSubmatch("peach punch")) // Submatch includes info about both the whole pattern matches and submatches within those matches

    fmt.Println(r.FindStringSubmatchIndex("peach punch")) // Returns info about indexes of matches and submatches

    fmt.Println(r.FindAllString("peach punch pinch", -1)) //All variants of these functions apply to all matches in the input, not just the first. 

    fmt.Println("all:", r.FindAllStringSubmatchIndex("peach punch pinch", -1))

    fmt.Println(r.FindAllString("peach punch pinch", 2)) //Providing a non negative integer as 2nd argument to these functions will limit the number of matches

    fmt.Println(r.Match([]byte("peach"))) //Can also provide []byte arguments

    r = regexp.MustCompile("p([a-z]+)ch") // This panics instead of returning the error which makes it safer to use for global variables
    fmt.Println("regexp:", r)

    fmt.Println(r.ReplaceAllString("a peach", "<fruit>")) //Used to replace subsets of strings with other values

    in := []byte("a peach")
    out := r.ReplaceAllFunc(in, bytes.ToUpper)

    fmt.Println(string(out))
}
