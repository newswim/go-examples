package main

import "fmt"

// Variadic function

func main () {

        fmt.Println(subtractThem(100,50, 25))

}

func subtractThem(args ...int) int {

        finalValue := 0

        for _, value := range args {

                finalValue -= value

        }

        return finalValue

}
