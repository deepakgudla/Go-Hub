package main

import "fmt"

type Operation_ struct{}

func (o Operation_) Name() string {
	return "Function"
}

// no params no returns
func name() {
	fmt.Println("type-1: no params and no returns")
}

// one param and no return
func nameTwo(a string) {
	fmt.Println("type-2:", a)
}

// one param and one return
func nameThree(b string) string {
	return fmt.Sprint("type-3:", b)
}

// two params and two return
func nameFive(typ string, _ string) (string, string) {
	no := "two return"
	return fmt.Sprint(typ, " two params and"), no
}

func nameFour(format string, _ int) (string, int) {
	Overs := 90
	return fmt.Sprint("no of overs played in a day in ", format), Overs
}

func (o Operation_) Run() {
	fmt.Println("---FUNCTIONS---")
	name()
	nameTwo("one param and no return")
	b := nameThree(" one param and one return")
	fmt.Println(b)

	typ, no := nameFive("type-4:", "")
	fmt.Println(typ, no)

	format, overs := nameFour("test cricket:", 90)
	fmt.Println(format, overs)

	func(z int) {
		fmt.Println("Anonymous number:", z)
	}(1357)
}

func init() {
	Register(Operation_{})
}

/* Anonymous function
syntax --> func(){}()
*/
