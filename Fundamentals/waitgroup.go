package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

type WG struct{}

func (w WG) Name() string {
	return "Wait_Group"
}

func (w WG) Run() {
	fmt.Println("goroutines\t", runtime.NumGoroutine())
	wg.Add(1)

	go abcd()
	mnop()

	fmt.Println("goroutines\t", runtime.NumGoroutine())

	wg.Wait()
}

func abcd() {
	for i := 0; i < 4; i++ {
		fmt.Println("foofoo", i)
	}
}

func mnop() {
	for i := 0; i < 5; i++ {
		fmt.Println("mnop-count", i)
	}
	wg.Done()
}

func init() {
	Register(WG{})
}
