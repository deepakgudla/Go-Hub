// package main

// import "fmt"

// func main() {
// 	a := factorial(13)
// 	fmt.Println("Total:", a)
// }

// func factorial (n int)  int {
// 	total := 1
// 	for i:=n; i>0; i-- {
// 		total *= i
// 	}

// 	return total
// }

package main

import (
	"fmt"
)

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&num)
	fmt.Printf("Factorial of %d is %d\n", num, factorial(num))
}