//converting string to number from input

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type number int

func (c number) String() string {
	return fmt.Sprint("number is : ", strconv.Itoa(int(c)))
}

func main() {
	//ascii to str conversion..
	var z number = 1357
	fmt.Println(z)

	fmt.Println("Hello World")
	fmt.Println("enter any number")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("number entered:", input)
	//parsefloat converts the string to a float number
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("added 1 to the entered number:", numRating+1)
	}

}
