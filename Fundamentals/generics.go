package main

import "fmt"

type Generic_ struct{}

func (g Generic_) Name() string {
	return "Generic"
}

func addA(a, b int) int {
	return a + b
}

func addB(a, b float64) float64 {
	return a + b
}

//[G int | float64] --> this is called as generic which accepts either of the parameter
func genericG[G int | float64](a, b G) G {
	return a + b
}

//instead of declaring data type in generic, we can do as such
type addNumbers interface {
	int | float64
}

//[G int | float64] --> this is called as generic which accepts either of the parameter
func genericG2[G addNumbers](a, b G) G {
	return a + b
}

func (g Generic_) Run() {
	fmt.Println(addA(1, 3))
	fmt.Println(addB(1.35, 0.77))
	fmt.Println("---generic func---")
	fmt.Println(genericG(1, 123))
	fmt.Println(genericG(1.35, 7.77))
	fmt.Println("type set....")
	fmt.Println("addition of numbers using type set interface :", genericG2(1, 1357))
}

func init() {
	Register(Generic_{})
}
