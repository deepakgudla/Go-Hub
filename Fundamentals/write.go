package main

import (
	"log"
	"os"
)

// creates a file
func main() {
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
