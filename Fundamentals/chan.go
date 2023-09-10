//sending and receiving in channels

package main

import (
	"fmt"
)

func main() {
	z := make(chan int)
	//z := make(chan int, 2) buffer channel

	go send(z) //send

	receive(z)

	fmt.Println("exit...")
}

func send(z chan<- int) {
	z <- 1
	z <- 3
}

func receive(z <-chan int) {
	fmt.Println("received value", <-z)
	// fmt.Println(<-z) printing buffer
}
