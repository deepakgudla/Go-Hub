package main

import (
	"fmt"
	"sort"
)

type Sort_ struct{}

func (s Sort_) Name() string {
	return "Sort"
}

func (s Sort_) Run() {
	a := []int{1, 77, 1357, 35, 5, 3, 7}

	fmt.Println("before sorting:", a)
	sort.Ints(a)
	fmt.Println("After sorting:", a)
}

func init() {
	Register(Sort_{})
}
