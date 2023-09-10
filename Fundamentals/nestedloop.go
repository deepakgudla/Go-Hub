package main 

import "fmt"

func main() {

	for i:=0; i<5; i++ { //this is outer loop
		fmt.Println("...")
		for j :=0; j<5; j++ { //this is inner loop
			fmt.Printf("outer loop value: %v & inner loop value: %v\n", i, j)
		}
	}
}
