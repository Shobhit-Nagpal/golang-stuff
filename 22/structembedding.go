package main

import "fmt"

type base struct {
    num int
}

func (b base) describe() string {
    return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
    base //container embeds a base
    str string
}

func main() {
    co := container{
        base: base{
            num: 1,
        },
        str: "deeznuts",
    }

    fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)
    fmt.Println("also num", co.base.num)
    fmt.Println("describe:", co.describe()) //As container embeds base, methods of base also become methods of containers

    type describer interface {
        describe() string
    }
    
    var d describer = co
    fmt.Println("describer:", d.describe()) //Embedding structs with methods may be used to add interface implementations on other structs
}

//Go supports embedding of structs and interfaces
