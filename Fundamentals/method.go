//attaching method to a type

package main

import "fmt"

type sport struct {
	name string
}

type Method_ struct{}

func (m Method_) Name() string {
	return "Method"
}

func (a sport) play() {
	fmt.Println("I play", a.name)
}

func (m Method_) Run() {
	b := sport{
		name: "cricket",
	}

	b.play()
}

func init() {
	Register(Method_{})
}
