//range clause in channels

package main

import "fmt"

func main() {
	z := make(chan int)

	go func() {
		for a:= 0; a<5;  a++ {
			z <-a
		}
		close(z)
	}()

	for m :=range z {
		fmt.Println(m)
	} 
}