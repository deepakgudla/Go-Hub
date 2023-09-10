package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 77, 1357, 35, 5, 3, 7}

	fmt.Println("before sorting:", a)
	sort.Ints(a)
	fmt.Println("After sorting:", a)
}