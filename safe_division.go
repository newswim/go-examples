package main

import "fmt"

func main () {

        fmt.Println(safeDiv(3, 0))
        fmt.Println(safeDiv(3, 2))

}

func safeDiv (num1, num2 int) int {

        defer func() {
                // "recover" lets us --catch--, if an error were to occur
                fmt.Println(recover())
        }()

        solution := num1 / num2

        return solution
}


// Notice that execution continues even though the first operation throws
// an error, this is thanks to the call to recover as the last step of safeDiv

// Printing the error is optional, simply remove the fmt.Println portion.
