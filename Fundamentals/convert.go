//converting string to number from input

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// type number int

type Convert struct{}

func (c Convert) Name() string {
	return "Convert"
}

func (c Convert) Run() {
	fmt.Print("Enter any number: ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Failed to read input:", err)
		return
	}

	input = strings.TrimSpace(input)
	num, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Printf("Invalid number %q: %v\n", input, err)
		return
	}

	fmt.Printf("Number entered: %.2f\n", num)
	fmt.Printf("After adding 1: %.2f\n", num+1)
}

func init() {
	Register(Convert{})
}
