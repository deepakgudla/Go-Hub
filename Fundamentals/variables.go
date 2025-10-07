package main

import "fmt"

const HelloWorld string = "hello-world"
const helloworld string = "qwertyuiop"

type Store struct{}

func (s Store) Name() string {
	return "Store"
}

func (s Store) Run() {
	//string datatype
	var username string = "golang"
	fmt.Println(username)
	//to know the datatpe
	fmt.Printf("datatype is of type: %T \n", username)

	//boolean datatype
	var isgolang bool = true
	fmt.Println(isgolang)
	fmt.Printf("bool is a datype datatype: %T \n", isgolang)

	//float datatype
	var decimal float32 = 1.123456789
	fmt.Println(decimal)

	var decimal_ float64 = 135.1234567890
	fmt.Println(decimal_)

	var testvar int
	fmt.Println("testing undeclared value for int....")
	fmt.Println(testvar)

	var testinteger string
	fmt.Println(testinteger)
	fmt.Printf("testvar is of type : %T \n", testinteger)
	//output for this func - empty string

	var lang = "go"
	fmt.Println(lang)

	fmt.Println(HelloWorld)
	fmt.Println(helloworld)

	//we can declare & initialize multiple var all at once
	a, b, c, d, e := 1, 2, 3, 4, 5
	fmt.Println(a, b, c, d, e)
}

func init() {
	Register(Store{})
}
