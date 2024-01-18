package main

import "fmt"

//ranges in slice
func main() {
	xy := []int{1, 3, 5, 7, 9, 10, 11, 12, 13, 14, 15}
	for y, a := range xy {
		fmt.Println("slice range", y, a)
	}

	z := map[string]int{
		"go":   2,
		"lang": 4,
	}
	for k, v := range z { //k,v = key value pair
		fmt.Println("map range:", k, v)
	}
}

/*
in place of xy & (y,a) we can consider any var
*/
