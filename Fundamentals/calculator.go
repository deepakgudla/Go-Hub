package main

import "fmt"

type Calculator struct{}

func (c Calculator) Name() string {
	return "Calculator"
}

func (c Calculator) Run() {
	fmt.Println("Running Calculator Example")
	defer fmt.Println()

	var a, b int
	fmt.Print("Enter two numbers (space separated): ")
	fmt.Scanln(&a, &b)

	fmt.Printf("a + b = %d\n", a+b)
	fmt.Printf("a - b = %d\n", a-b)
	fmt.Printf("a * b = %d\n", a*b)

	if b != 0 {
		fmt.Printf("a / b = %d\n", a/b)
		fmt.Printf("a %% b = %d\n", a%b)
	} else {
		fmt.Println("Cannot divide by zero")
	}
}

// Auto-register on package init
func init() {
	Register(Calculator{})
}
