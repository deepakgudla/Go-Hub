package main

import (
	"fmt"
	"log"
	"net/http"
)

type StartServer struct{}

func (s StartServer) Name() string {
	return "Start_Server"
}

type Cricket struct {
	Format string `json:"format"`
	Overs  string `json:"overs"`
}

// func cricket_(a http.ResponseWriter, b *http.Request) {
// 	fmt.Println("cricket is a lovely sport")
// }

func cricketHandler(a http.ResponseWriter, b *http.Request) {
	fmt.Fprintf(a, "cricket is a lovely sport")
}

func handleRequest() {
	http.HandleFunc("/", cricketHandler)
	log.Fatal(http.ListenAndServe(":1357", nil))
}

func (s StartServer) Run() {
	handleRequest()
}

func init() {
	Register(StartServer{})
}
