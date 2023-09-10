package main

import "fmt"

//passing by value

func intPtr(c *int) {
	*c = 13
}

func slicePtr(m []int) {
	m[0] = 13
}

func mapPtr(y map[string]int, z string) {
	y[z] = 17
}

func main() {
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
	fmt.Println("pass by value : ", a)
	//slices-pointers.. pass by value
	fmt.Println(".....slices & pointers.....")
	n := []int{1, 3, 5, 7}
	fmt.Println(n)
	slicePtr(n)
	fmt.Println(n)

	//maps-pointers - pass by value
	x := make(map[string]int)
	x["golang"] = 77
	fmt.Println(x["golang"])
	mapPtr(x, "go")
	fmt.Println(x["go"])

}
