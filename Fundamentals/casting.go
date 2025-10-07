package main

import "fmt"

type TypeCast struct{}

func (tc TypeCast) Name() string {
	return "TypeCast"
}

func (tc TypeCast) Run() {

	x := 1357
	y := 13.57
	fmt.Printf("%v is of type %T\n", x, x)
	//the above statement gives the output 1357 & int
	fmt.Printf("%v is of type %T\n", y, y)
	//there is no concept of typecasting in the above case...

	// In golang we cannot take a value that is float32 and store  it in a var that is declared to hold a value of float64

	//typecasting does not work in the below case; reason being the above statement
	var z float32 = 1.2345
	fmt.Printf("%v is of type %T\n", z, z)

	y = float64(z)
	fmt.Printf("%v is of type %T\n", y, y)
}

func init() {
	Register(TypeCast{})
}

// type casting can be done if the both the values are of same data type
