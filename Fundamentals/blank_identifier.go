package main

import "fmt"

type BlankIdentifier struct{}

func (b BlankIdentifier) Name() string {
	return "BlankIdentifier"
}

func (b BlankIdentifier) Run() {
	const name = "qwerty"
	const age = 100
	fmt.Printf("%s is %d years old\n", name, age)

	e, c, _, f := 0, 1, 2, "golang"
	fmt.Println(e, c, f)
}

func init() {
	Register(BlankIdentifier{})
}

//	" _ " is the blank identifier which tells Go to ignore that value.
