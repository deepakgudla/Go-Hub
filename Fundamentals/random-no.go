package main

import (
	"fmt"
	"math/rand"
)

type RandomNumberGenerator struct{}

func (r RandomNumberGenerator) Name() string {
	return "RandomNumberGenerator"
}

func (r RandomNumberGenerator) Run() {
	fmt.Println("The number is:", rand.Intn(999))
}

func init() {
	Register(RandomNumberGenerator{})
}

//internal working ?
