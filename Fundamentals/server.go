package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type About struct {
	Language  string `json:"language"`
	Framework string `json:"framework"`
}

type Server_ struct{}

func (s Server_) Name() string {
	return "Server"
}

func (s Server_) Run() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)

	// Start the server
	fmt.Println("Server listening on port 2468...")
	log.Fatal(http.ListenAndServe(":2468", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	// response body
	fmt.Fprintf(w, "<h1> GOLANG SERVER RESPONSE </h1>")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	about := About{
		Language:  "javascript  js",
		Framework: "reactjs(frontend framework)",
	}
	w.Header().Set("Content-Type", "application/json")

	jsonData, err := json.Marshal(about)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(jsonData); err != nil {
		http.Error(w, "failed to write response", http.StatusInternalServerError)
		return
	}

}

func init() {
	Register(Server_{})
}
