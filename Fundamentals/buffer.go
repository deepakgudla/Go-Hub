// buffer
package main

import (
	"bytes"
	"fmt"
)

type Buffer struct{}

func (b Buffer) Name() string {
	return "Buffer"
}

func (b Buffer) Run() {
	fmt.Println("buffers")
	a := bytes.NewBufferString("Test")
	a.WriteString("cricket")
	fmt.Println(a.String())

	//can rewrite the string using reset
	a.Reset()
	a.WriteString("badminton")
	fmt.Println(a.String())

	//writing
	a.Write([]byte("buffer"))
	fmt.Println(a.String())
}

func init() {
	Register(Buffer{})
}
