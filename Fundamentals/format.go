package main

import "fmt"

func main () {
	x := 1
	y := 3
	var language string = "golang"

	fmt.Println("**format verbs**")
	fmt.Printf("The value of x is %v and its data type is %T\n", x,x)
	fmt.Printf("value of y is %d\n", y)
	fmt.Printf("This  is a %s language\n", language)
}