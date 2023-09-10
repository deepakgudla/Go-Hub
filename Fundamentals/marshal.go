package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name string
	Grade string
	Gpa int
}

type exam struct {}

func main( ) {
	fmt.Println("marshal aka serialization")

	c1 := course{
		Name: "computer networks",
		Grade: "B",
		Gpa: 8,
	}

	c2 := course {
		Name: "Blockchain",
		Grade: "A",
		Gpa: 10,
	}

	student := []course{c1, c2}
	fmt.Println(student)

	ab, err := json.Marshal(student)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(ab))
}