package main

import "fmt"

func main () {
	x := 1

	switch {
	case x < 1:
		fmt.Println("value is less than 1")
	case x == 1:
		fmt.Println("value is equal to 1")
	default:
		fmt.Println("This is the default case")
	}
	fmt.Println("---------------------")

	switch x {
	case 0:
		fmt.Println("x is not 0")
		fallthrough
	case 1:
		fmt.Println("fallthrough x is 1 but no assignment")
	case 2:
		fmt.Println("print sth")
	default:
		fmt.Println("default case for x") 
	}
}

// no break statement for switch....?

