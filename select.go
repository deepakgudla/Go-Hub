//lecture: 69, should relearn...

package main

import "fmt"

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	select {
	case a := <-ch1:
		fmt.Println("value from channel1", a)
	case b := <-ch2:
		fmt.Println("value from channel2", b)
	}
}
