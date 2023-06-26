package main

import "fmt"

func sum(a ...int) int { //unlimited set of arguments... ( ... )
	fmt.Println(a)

	i := 0
	for _, v := range a {
		i += v
	}

	return i
}

func main() {
	fmt.Println("---variadic parameter")
	b := sum(1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 777, 1357)
	fmt.Println("sum of all the values", b)

}
