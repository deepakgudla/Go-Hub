package main

import "fmt"

type IE_ struct{}

func (i IE_) Name() string {
	return "If_else"
}

func (ie IE_) Run() {
	x := 1
	if x > 0 {
		fmt.Printf("the value of x is %v\n", x)
	} else {
		fmt.Println("else statement")
	}
}

func init() {
	Register(IE_{})
}
