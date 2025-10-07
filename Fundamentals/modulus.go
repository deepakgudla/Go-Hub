package main

import "fmt"

type Mod struct{}

func (m Mod) Name() string {
	return "Mod"
}

func (m Mod) Run() {

	a := 22 / 10
	b := 22 % 10
	fmt.Println(a, b)

	for i := 0; i < 20; i++ {
		if i%2 != 0 {
			fmt.Printf("odd number is %v\n", i)
		}
	}
}

func init() {
	Register(Mod{})
}
