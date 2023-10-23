package main

import (
    "fmt"
    fp "path/filepath"
    "strings"
)

func main() {
    p := fp.Join("dir1", "dir2", "filename") //Join should be used to construct paths in a portable way
    fmt.Println("p:", p)

    fmt.Println(fp.Join("dir1//", "filename"))
    fmt.Println(fp.Join("dir1/../dir1", "filename"))

    fmt.Println("Dir(p):", fp.Dir(p))
    fmt.Println("Base(p):", fp.Base(p))

    fmt.Println(fp.IsAbs("dir/file"))
    fmt.Println(fp.IsAbs("/dir/file"))

    filename := "config.json"
    ext := fp.Ext(filename)
    fmt.Println(ext)

    fmt.Println(strings.TrimSuffix(filename, ext)) //To get the filename with the extension removed

    rel, err := fp.Rel("a/b", "a/b/t/file") //Rel finds relative path b/w base and target. Returns error if target cannot be made relative to base
    if err != nil {
        panic(err)
    }
    
    fmt.Println(rel)

    rel, err = fp.Rel("a/b", "a/c/t/file")
    if err != nil {
        panic(err)
    }
    fmt.Println(rel)
}
