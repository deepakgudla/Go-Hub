package main

import (
	"fmt"
	"time"
)

type Timestamp struct{}

func (t Timestamp) Name() string {
	return "Timestamp"
}

func (t Timestamp) Run() {

	fmt.Println("Current Time:")

	currentTime := time.Now()
	fmt.Println(currentTime)
	fmt.Println(currentTime.Format("01-02-2006 17:04:05 Monday"))

	// createdDate := time.Date(2001, time.January)
	// fmt.Println(createdDate)
}

func init() {
	Register(Timestamp{})
}
