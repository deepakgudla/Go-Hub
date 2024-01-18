// buffer
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println("buffers")
	a := bytes.NewBufferString("Test")
	a.WriteString("cricket")
	fmt.Println(a.String())
	//can rewrite the string using reset
	a.Reset()
	a.WriteString("badminton")
	fmt.Println(a.String())

	//writing
	a.Write([]byte("byte buffer"))
	fmt.Println(a.String())
}
