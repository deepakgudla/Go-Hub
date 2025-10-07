package main

import "fmt"

type cricket struct {
	format string
}

type game interface {
	play()
	franchise()
}

type SetMethod_ struct{}

func (s SetMethod_) Name() string {
	return "Set_Method"
}

func (b cricket) play() {
	fmt.Println("ultimate format of cricket is", b.format)
}

func (b *cricket) franchise() {
	b.format = "ipl"
	fmt.Println("RCB won", b.format, "Trophy...")
}

func gameCricket(z game) {
	z.play()
}

func (s SetMethod_) Run() {

	//value semantic of type T with receiver T
	a := cricket{"test"}
	a.play()
	//gameCricket(a) - throws an error since method franchise has ptr receiver

	//Pointer semantic of type *T with receiver *T and T
	c := &cricket{"odi"}
	//c.play()
	c.franchise()
	fmt.Println("..")
	gameCricket(c)

	/*
		cricket does not implement game (method franchise has pointer receiver)
	*/
}

func init() {
	Register(SetMethod_{})
}
