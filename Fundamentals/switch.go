package main

import "fmt"

type Switch_Loop struct{}

func (s Switch_Loop) Name() string {
	return "Switch"
}

func (s Switch_Loop) Run() {
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

func init() {
	Register(Switch_Loop{})
}

// no break statement for switch....?
