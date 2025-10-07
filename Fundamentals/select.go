//lecture: 69, should relearn...

package main

import "fmt"

type Select_ struct{}

func (s Select_) Name() string {
	return "Select"
}

func (s Select_) Run() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		ch1 <- 10
	}()

	go func() {
		ch2 <- 20
	}()

	select {
	case a := <-ch1:
		fmt.Println("value from ch1", a)
	case b := <-ch2:
		fmt.Println("value from ch2", b)
	}
}

func init() {
	Register(Select_{})
}
