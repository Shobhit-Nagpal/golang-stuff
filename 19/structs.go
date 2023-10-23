package main

import "fmt"

type person struct {
    name string
    age int
}

func newPerson(name string) *person {
    p := person{name: name}
    p.age = 42
    return &p
}

func main() {
    fmt.Println(person{"Bob", 20})
    fmt.Println(person{name: "Alice", age: 30})
    fmt.Println(person{name: "Fred"})
    fmt.Println(&person{name: "Ann", age: 40}) //Gives a pointer to this struct
    fmt.Println(newPerson("Jon"))

    s := person{name: "Sean", age: 50}
    fmt.Println(s.name)

    sp := &s
    fmt.Println(sp.age) // Can use dots with struct pointers, the pointers are automatically dereferenced

    sp.age = 51
    fmt.Println(sp.age)

    dog := struct {
        name string
        isGood bool
    }{
        "Rex",
        true,
    } //If we use struct only for a single value, we don't need to give it a name, the value can have an anonymous struct type
    fmt.Println(dog)
}

//Go structs are a typed collection of fields
//Structs are mutable
