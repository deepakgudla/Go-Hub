package main

import (
	"fmt"
)

func main() {
	a := make(chan int, 1) //method-2 adding buffer
	// values are proportional to buffer 2 values the buffer value is 2
	// b := make(chan int)
	m := make(<-chan int) //receive
	n := make(chan<- int) //send

	a <- 13
	a <- 11
	fmt.Println(<-a) //channel block..

	fmt.Printf("m\t%T\n", m)

	//method-1 of avoiding channel block
	// go func() {
	// 	b <- 11
	// }()

	// fmt.Println(<-b)
}
