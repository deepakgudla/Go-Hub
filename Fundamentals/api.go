package main

import (
	"fmt"
	"log"
	"net/http"
)

type Cricket struct {
	Format string `json:"format"`
	Overs  string `json:"overs"`
}

type games []Cricket

func cricket_(a http.ResponseWriter, b *http.Request) {
	fmt.Println("cricket is a lovely sport")
}

func abcd(a http.ResponseWriter, b *http.Request) {
	fmt.Fprintf(a, "cricket is a lovely sport")
}

func handleRequest() {
	http.HandleFunc("/", abcd)
	log.Fatal(http.ListenAndServe(":1357", nil))
}

func main() {
	handleRequest()
}
