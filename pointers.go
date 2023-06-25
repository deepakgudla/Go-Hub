package main

import "fmt"

//passing by value

func intPtr(c *int) {
	*c = 13
}


func main () {
	fmt.Println("-----pointers-----")
	a := 1357
	var sport string = "cricket"
	b := &a
	fmt.Println(a)
	//& - returns the address of the var
	fmt.Printf("address of the var %v is  %v \n", a, &a) //printing the address of the var a
	fmt.Printf("data type of %v is %T and its address is %v\n", sport, &sport, &sport)
	fmt.Printf("address of the var %v is  %v \n", b, &b)
	fmt.Printf("value is %v\n", *&a)

	*b = 777
	fmt.Println(a)

	//prinitng pass by value..
	intPtr(&a)
	fmt.Print("pass by value : ", a)

}