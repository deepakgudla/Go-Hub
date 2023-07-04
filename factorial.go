package main

import "fmt"

func main() {
	a := factorial(13)
	fmt.Println("Total:", a)
}

func factorial (n int)  int {
	total := 1
	for i:=n; i>0; i-- {
		total *= i
	}

	return total
}