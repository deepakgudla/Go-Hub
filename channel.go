package main

import (
	"fmt"
)

func main() {
	fmt.Println("---channels---")

	a := make (chan int)

	// go func () {
	// 	for i:=0;  i<3; i++ {
	// 		a <-i
	// 	}
	// }()

	// go func() {
	// 	for {
	// 		fmt.Println(<-a)
	// 	}
	// } ()

	// time.Sleep(time.Minute)

	//---------------------------
	go abc(a) //sending

	mnop(a)

	fmt.Println(".......exit.......")

}

func abc(a chan<- int) {
	a <- 1357
}

func mnop(a<- chan int) {
	fmt.Println(<-a)
}

