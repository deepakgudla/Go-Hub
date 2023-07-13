//using arguments in channel

package main

import "fmt"

func main() {
	a := incrementor()
	aSum := puller(a)
	for n := range aSum {
		fmt.Println(n)
	} 
}

func incrementor() chan int {
	out := make(chan int)
	go func() {
		for i := 0; i <10; i++ {
			out <-i
		}
		close(out)
	}()
	return out
}

func puller(a chan int) chan int {
	out := make(chan int)
	go func() {
		var sum int
		for n := range a {
			sum += n
		}
		out <-sum
		close(out)
	}()
	return out
}