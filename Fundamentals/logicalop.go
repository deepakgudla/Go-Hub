package main

import "fmt"

type LogicalOp struct{}

func (l LogicalOp) Name() string {
	return "Logical_Operators"
}

func (l LogicalOp) Run() {
	x := 1
	y := 2

	if x > 0 && y > 1 {
		fmt.Printf("%v is less and %v is greater\n", x, y)
	}

	if x > 0 && y < 3 {
		fmt.Println("met the condition")
	}

	if x == 1 || y < 3 {
		fmt.Printf("only the %v met the condition\n", x)
	}
}

func init() {
	Register(LogicalOp{})
}
