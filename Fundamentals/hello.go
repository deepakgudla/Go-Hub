package main

import "fmt"

type Greetings struct{}

func (g Greetings) Name() string {
	return "HelloWorld"
}

func (g Greetings) Run() {
	fmt.Println("Hello, Go")
}

func init() {
	Register(Greetings{})
}
