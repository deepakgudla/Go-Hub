package main

import "fmt"

func main() {
	a := 0x38f
	//to print a number as binary add (%b) tag...
	fmt.Printf("8 as binary %b\n", a)
	fmt.Printf("")
	//to print a number in hexadecial use this tag --> (%x)
	fmt.Printf("1357 as hexadecimal %x\n", a)
	//default binary and hexadecimal

	fmt.Printf("%v \t %b \t %X\n", a, a, a)
	//to print hexadecimal values with 0x
	fmt.Printf("%v \t %b \t %#X\n", a, a, a)
	fmt.Printf("the value %v is of type %T\n", a, a)

}
