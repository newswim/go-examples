package main

import "fmt"


func main () {

        fmt.Println(factorial(5))

}

func factorial (num int) int {

        // we need a way to end the recursion

        if num == 0 {

                return 1
        }

        return num * factorial(num - 1)
}
