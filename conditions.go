package main

import "fmt"

func main () {
	x := 4
	y := 135

	if x > 3 {
		fmt.Println("number is greater than 3")
	} else {
		fmt.Println("entered number is less than 3")
	}

	if y < 100 {
		fmt.Println("number is less than 135") 
	} else if  y == 135 {
		fmt.Println("y is equal to 135")
	} else {
		fmt.Println("print sth")
	}




}

//incorrect syntax 
/*
if ... {
	fmt.Println("....")
}
//the following is not the correct syntax
// else should be adjacent to the closing bracket of if...
else {
}

*/
