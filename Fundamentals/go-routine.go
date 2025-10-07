//go-routine example

package main

import (
	"fmt"
	"time"
)

type GoRoutine_ struct{}

func (g GoRoutine_) Name() string {
	return "GoRoutine"
}

func task(task string) {
	fmt.Println("doing task-1", task)
	time.Sleep(1 * time.Second)
	fmt.Println("doing task-2", task)
}

func (g GoRoutine_) Run() {
	go task("drinking coffee")
	task("writing journal")
}

func init() {
	Register(GoRoutine_{})
}
