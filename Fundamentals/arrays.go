package main

import "fmt"

func main() {
        fmt.Println("-------arrays-------")
        //method #1 of declaring array
        var x[5] int
        fmt.Println(x) //prints an array of size 5 with the value "0"
        x [1] = 4
        fmt.Println(x) //prints 4 in the designated position
        fmt.Println(len(x))

        //method #2 of declaring array
        //array literal compiler count elements
        z := [...]string{"test", "odi"}
        fmt.Printf("%#v and the datatype is %T\n", z,z)

        //method #3 of declaring array
        //array literal
        y := [5]int{1, 2, 3, 4, 5}
        fmt.Printf("%#v and the datatype is %T\n", y,y)

        fmt.Println(len(z))
        fmt.Println(len(y))

}
