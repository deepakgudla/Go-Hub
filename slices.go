package main

import "fmt"

func main() {
        var sports = []string{"cricket", "badminton", "tennis"}
        fmt.Println(sports)
        
        //adding elements to the existing slice
        sports = append(sports, "football", "volleyball", "basketball", "golf", "wrestling")
        fmt.Println(sports)
        sports = append(sports[1:5]) //prints 1 2 3 4 
        fmt.Println(sports)
        sports = append(sports[2:6])// prints 2 3 4 5
        fmt.Println(sports)
        //sports = append(sports[1:5][2:6]) // ?
        //fmt.Println(sports)
        fmt.Println("_________________")

        a := []string{"hello", "gopher"}
        fmt.Println(a)

        for i,v := range(sports) {
                fmt.Printf("index %v is storing the value %v\n", i,v)
        }
}
