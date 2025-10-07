package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type PasswordCheck struct{}

func (p PasswordCheck) Name() string {
	return "PasswordCheck"

}

func (p PasswordCheck) Run() {
	a := `qwertymnop`
	ab, err := bcrypt.GenerateFromPassword([]byte(a), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}

	//fmt.Println(ab)
	pwd := `qwertymnop`

	err = bcrypt.CompareHashAndPassword(ab, []byte(pwd)) //compares the hash & pwd
	if err != nil {
		fmt.Println("wrong password, please try again!")
		return
	}
	fmt.Println("password matched...........")
}

func init() {
	Register(PasswordCheck{})
}
