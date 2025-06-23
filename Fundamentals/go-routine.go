//go-routine example

package main

import (
	"fmt"
	"time"
)

func task(task string) {
	fmt.Println("doing task-1", task)
	time.Sleep(1 * time.Second)
	fmt.Println("doing task-2", task)
}

func main() {
	go task("drinking coffee")
	task("writing journal")
}
