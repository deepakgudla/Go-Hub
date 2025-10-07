package main

import "fmt"

type NestedLoop struct{}

func (n NestedLoop) Name() string {
	return "Nestedloop"
}

func (n NestedLoop) Run() {

	for i := 0; i < 5; i++ { //this is outer loop
		fmt.Println("...")
		for j := 0; j < 5; j++ { //this is inner loop
			fmt.Printf("outer loop value: %v & inner loop value: %v\n", i, j)
		}
	}
}

func init() {
	Register(NestedLoop{})
}
