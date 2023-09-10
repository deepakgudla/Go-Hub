//attaching method to a type

package main

import "fmt"

type sport struct {
	name string
}

func (a sport) play() {
	fmt.Println("I play", a.name)
}

func main() {
	b := sport{
		name: "cricket",
	}

	b.play()
}
