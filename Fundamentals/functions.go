package main

import "fmt"

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
func nameFive(typ string, no string) (string, string) {
	no = "two return"
	return fmt.Sprint(typ, " two params and"), no
}

func nameFour(format string, overs int) (string, int) {
	overs = 90
	return fmt.Sprint("no of overs played in a day in ", format), overs
}

func main() {
	fmt.Println("---FUNCTIONS---")
	name()
	nameTwo("one param and no return")
	b := nameThree(" one param and one return")
	fmt.Println(b)

	typ, no := nameFive("type-4:", "")
	fmt.Println(typ, no)

	format, overs := nameFour("test cricket:", 90)
	fmt.Println(format, overs)
	/* Anonymous function
	syntax --> func(){}()
	*/
	func(z int) {
		fmt.Println("Anonymous number is:", z)
	}(1357)
}
