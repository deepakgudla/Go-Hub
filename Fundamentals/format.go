/*
%s is for strings
%d for integers
%U for unicode numbers...
*/

package main

import "fmt"

type FormatVerbs struct{}

func (f FormatVerbs) Name() string {
	return "Format_Verbs"
}

func (f FormatVerbs) Run() {
	x := 1
	y := 3
	var language string = "golang"

	fmt.Println("**format verbs**")
	fmt.Printf("The value of x is %v and its data type is %T\n", x, x)
	fmt.Printf("value of y is %d\n", y)
	fmt.Printf("This  is a %s language\n", language)

	const name = " Criket ODI worldcup"
	const period = 4
	//unicode := 2693
	fmt.Printf("%s is played for every %d years.\n", name, period)
	//fmt.Println("Unicode number: %u\n", unicode);
}

func init() {
	Register(FormatVerbs{})
}
