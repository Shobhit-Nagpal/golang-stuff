package main

import "fmt"

type rect struct {
    width, height int
}

func (r *rect) area() int { //This method has a reciever type of *rect
    return r.width * r.height
}

func (r rect) perim() int { // Methods can be defined for either pointer or value reciever types
    return 2*r.width + 2*r.height
}

func main() {
    r := rect{width: 10, height: 5}

    fmt.Println("Area:", r.area()) //Calling the methods
    fmt.Println("Perimeter:", r.perim())

    rp := &r
    fmt.Println("Area:", rp.area())
    fmt.Println("Perimeter:", rp.perim())

}

//Go automatically handles the conversion between pointers and values for method calls. We may want to use pointer receiver types to avoid copying on method calls or allow the method to mutate the struct 
