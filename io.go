//i/o

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	welcome := "helloworld"
	fmt.Println(welcome)

	/* to get input from the user
	we need to make use of bufio package
	*/

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter your name")

	//for storing the input
	// ~ to try catch we can use sth like [input, _/err]

	input, _ := reader.ReadString('\n')
	//fmt.Println(err)
	fmt.Print("name successfully registered as = ", input)
	//fmt.Println(err)
}
