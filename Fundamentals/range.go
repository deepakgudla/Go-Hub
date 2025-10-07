//range clause in channels

package main

import "fmt"

type Interval struct{}

func (i Interval) Name() string {
	return "Range"
}

func (i Interval) Run() {
	z := make(chan int)

	go func() {
		for a := 0; a < 7; a++ {
			z <- a
		}
		close(z)
	}()

	for m := range z {
		fmt.Println(m)
	}
}

func init() {
	Register(Interval{})
}
