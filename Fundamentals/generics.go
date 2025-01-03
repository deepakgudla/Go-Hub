package main

import "fmt"

func addA(a, b int) int {
	return a + b
}

func addB(a, b float64) float64 {
	return a + b
}

func genericG[G int | float64](a, b G) G { //[G int | float64] --> this is called as generic which accepts either of the parameter
	return a + b
}

//instead of declaring data type in generic, we can do as such

type addNumbers interface {
	int | float64
}

func genericG2[G addNumbers](a, b G) G { //[G int | float64] --> this is called as generic which accepts either of the parameter
	return a + b
}

func main() {
	fmt.Println(addA(1, 3))
	fmt.Println(addB(1.35, 0.77))
	fmt.Println("---generic func---")
	fmt.Println(genericG[int](1, 123))
	fmt.Println(genericG[float64](1.35, 7.77))
	fmt.Println("type set....")
	fmt.Println("addition of numbers using type set interface :", genericG2[int](1, 1357))
}
