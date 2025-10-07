package main

import (
	"log"
	"os"
)

type WriteIntoFile struct{}

func (w WriteIntoFile) Name() string {
	return "WriteIntoFile"
}

// creates a file
func (w WriteIntoFile) Run() {
	f, err := os.Create("qwerty.txt")
	if err != nil {
		log.Fatalf("error %s", err)
	}

	defer f.Close()
	//writes below text to the file
	b := []byte("cricket & badminton")
	_, err = f.Write(b)
	if err != nil {
		log.Fatalf("error %s", err)
	}
}

func init() {
	Register(WriteIntoFile{})
}
