package main

import "fmt"

// "defer" will always run at the very end

func main () {

        defer printTwo()
        printOne()

}

func printOne() { fmt.Println(1) }
func printTwo() { fmt.Println(2) }
