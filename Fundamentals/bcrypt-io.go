package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type PasswordCheckIO struct{}

func (pc PasswordCheckIO) Name() string {
	return "PasswordCheck_IO"
}

func (pc PasswordCheckIO) Run() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your password 🔑: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("please confirm your password 🔍: ")
	confirmPassword, _ := reader.ReadString('\n')
	confirmPassword = strings.TrimSpace(confirmPassword)

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(confirmPassword))
	if err != nil {
		fmt.Println("password didn't matched ⚠")
		return
	}

	fmt.Println("Passwords matched ✅")
}

func init() {
	Register(PasswordCheckIO{})
}
