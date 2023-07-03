package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("---channels---")

	a := make (chan int)

	go func () {
		for i:=0;  i<3; i++ {
			a <-i
		}
	}()

	go func() {
		for {
			fmt.Println(<-a)
		}
	} ()

	time.Sleep(time.Minute)
}