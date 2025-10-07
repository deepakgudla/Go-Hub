package main

import "fmt"

type Loop_ struct{}

func (l Loop_) Name() string {
	return "For_loop"
}

func (l Loop_) Run() {
	x := 35
	y := 5
	fmt.Printf("value of x=%v, & y=%v\n", x, y)

	for i := 0; i < 5; i++ {
		fmt.Printf("counting till five %v first for loop\n", i)
	}

	//infinite loop
	for y < 10 {
		fmt.Printf("value of y is %v, loop 2 \n", y)
		y++ //with this condition its not infinite loop since y 		    is declared already
	}

	/*
		//  break--> using break to get out of this loop,
		//  using break like we are done running in this loop...
		//  if y is >20 it will stop if not it keeps on incrementing until it reaches 20,
		//  and then halts... it will print till 21 because 21 is > 20,
		//  and then it gets out of the loop
	*/

	for {
		fmt.Printf("value of y is %v, loop 3\n", y)
		if y > 20 {
			break
		}
		y++
	}

	/*
		continue --> step out of the current iteration loop,
		and go back to the next iteration of the loop
	*/

	for i := 0; i < 20; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println("even numbers are: ", i)
	}
}

func init() {
	Register(Loop_{})
}
