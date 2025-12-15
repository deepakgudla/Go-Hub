package main

import "fmt"

type FactorialExample struct{}

func (f FactorialExample) Name() string {
	return "Factorial"
}

// factorial logic
func (f FactorialExample) Run() {
	fmt.Println("Running Factorial Example")
	defer fmt.Println()

	var num int
	fmt.Print("Enter a number: ")
	if _, err := fmt.Scanln(&num); err != nil {
		fmt.Println("Invalid input", err)
	}
	fmt.Printf("Factorial of %d is %d\n", num, factorial(num))
}

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// Execution code
func init() {
	Register(FactorialExample{})
}
