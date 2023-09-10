package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("time in go program")

	currentTime := time.Now()
	fmt.Println(currentTime)
	fmt.Println(currentTime.Format("01-02-2006 17:04:05 Monday"))

	// createdDate := time.Date(2001, time.January)
	// fmt.Println(createdDate)
}
