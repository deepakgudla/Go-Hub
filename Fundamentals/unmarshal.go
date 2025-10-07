package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name  string `json:"Name"`
	Grade string `json:"Grade"`
	Gpa   int    `json:"Gpa"`
}

type Unmarsh struct{}

func (u Unmarsh) Name() string {
	return "Unmarsh"
}

func (u Unmarsh) Run() {

	a := `[{"Name":"computer networks","Grade":"B","Gpa":8},{"Name":"Blockchain","Grade":"A","Gpa":10}]`
	aa := []byte(a)
	var student []course
	err := json.Unmarshal(aa, &student)

	if err != nil {
		fmt.Println(err)
	}

	//fmt.Println("result: ", student)

	for i, v := range student {
		fmt.Println("student number :", i)
		fmt.Println(v.Name, v.Grade, v.Gpa)
	}
}

func init() {
	Register(Unmarsh{})
}

//output from "marshal.go"
// [{"Name":"computer networks","Grade":"B","Gpa":8},{"Name":"Blockchain","Grade":"A","Gpa":10}]
