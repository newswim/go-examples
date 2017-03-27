package closures

import "fmt"


func main () {

        num3 := 3

        doubleFunc := func() int {

                num3 *= 2

                return num3
        }

        fmt.Println(doubleFunc())
        fmt.Println(doubleFunc())

}
