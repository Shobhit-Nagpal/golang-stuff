package main

import "fmt"

func sum(nums ...int) {
    fmt.Print(nums, " ")
    total := 0

    //Within the fuction, the type of nums is []int
    for _, num := range nums {
        total += num
    }

    fmt.Println(total)
}

func main() {
    sum(1,2)
    sum(1,2,3)

    //If we have multiple arguments in slice, we can call the function the following way
    nums := []int{1,2,3,4}
    sum(nums...)
}

//Variadic functions can be called with any number of trailing arguments
