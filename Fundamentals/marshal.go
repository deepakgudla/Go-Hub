package main

import (
	"encoding/json"
	"fmt"
)

type graduate struct {
	Name  string
	Grade string
	Gpa   int
}

type Marsh_ struct{}

func (m Marsh_) Name() string {
	return "Marshal"
}

func (m Marsh_) Run() {
	fmt.Println("marshal aka serialization")

	c1 := graduate{
		Name:  "computer networks",
		Grade: "B",
		Gpa:   8,
	}

	c2 := graduate{
		Name:  "Blockchain",
		Grade: "A",
		Gpa:   10,
	}

	student := []graduate{c1, c2}
	fmt.Println(student)

	ab, err := json.Marshal(student)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(ab))
}

func init() {
	Register(Marsh_{})
}
