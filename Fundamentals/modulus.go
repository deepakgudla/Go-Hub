package main

import "fmt"

func main() {

	a := 22 / 10
	b := 22 % 10
	fmt.Println(a, b)

	for i :=0; i <20; i++ {
		 if i%2 != 0 {
			 fmt.Printf("odd number is %v\n", i)
		 }
	 }

}
