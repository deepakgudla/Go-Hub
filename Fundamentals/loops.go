package main

import "fmt"

func main() {
	x := 1
	if x > 0 {
		fmt.Printf("the value of x is %v\n", x)
	} else {
		fmt.Println("else statement")
	}
}