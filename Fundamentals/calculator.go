package main

import "fmt"

func main() {
	var x, y int

	fmt.Print("Enter the first number: ")
	fmt.Scanln(&x)

	fmt.Print("Enter the second number: ")
	fmt.Scanln(&y)

	sum := x + y
	fmt.Printf("Sum: %d\n", sum)

	diff := x - y
	fmt.Printf("Difference: %d\n", diff)

	product := x * y
	fmt.Printf("Product: %d\n", product)

	quotient := x / y
	fmt.Printf("Quotient: %d\n", quotient)

	remainder := x % y
	fmt.Printf("Remainder: %d\n", remainder)
}
